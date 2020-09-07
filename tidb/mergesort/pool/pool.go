package pool

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	QSize   int
	Workers int
	MaxIdle time.Duration
}

type Task interface {
	Run() error
	RunError(err error)
}

type Pool struct {
	Config
	waitingChan chan Task
	workersChan chan struct{}
	wg          sync.WaitGroup
	close       int32
}

func NewPool(cfg *Config) *Pool {

	return &Pool{
		Config:       *cfg,
		waitingChan:  make(chan Task, cfg.QSize),
		workersChan: make(chan struct{}, cfg.Workers),
		wg:           sync.WaitGroup{},
		close:        0,
	}
}

func (p *Pool) Put(t Task) *Pool {
	select {
	case p.workersChan <- struct{}{}:
		p.wg.Add(1)
		go p.work(t)
	default:
		select {
		case p.workersChan <- struct{}{}:
			p.wg.Add(1)
			go p.work(t)
		case p.waitingChan <- t:
		}
	}
	return p
}

func (p *Pool) work(t Task) {
	timer := time.NewTimer(p.MaxIdle)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Println("goroutine pool:", err)
			default:

			}
		}
	}()

	defer func() {
		timer.Stop()
		p.wg.Done()
		<- p.workersChan
	}()
	for {
		if err := t.Run(); err != nil {
			t.RunError(err)

		}
		select {
		case <-timer.C:
			return
		case newTask, ok := <-p.waitingChan:
			if !ok {
				return
			}
			if newTask == nil {
				return
			}

			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(p.MaxIdle)
			t = newTask
		}
	}
}

func (p *Pool) Close(grace bool) {
	if ok := atomic.CompareAndSwapInt32(&p.close, 0, 1); ok {
		close(p.workersChan)
		close(p.workersChan)
		if grace {
			p.wg.Wait()
		}
	}
}