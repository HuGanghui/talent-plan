package task

import (
	"pingcap/talentplan/tidb/mergesort/algorithm"
	"pingcap/talentplan/tidb/mergesort/pool"
	"runtime"
)

const SORTING_ARRAY_THRESHOLD = 1 << 12

type Sorter interface {
	Run()
}

type SingleSorter struct {
	sortingArray []int64
}

func (s *SingleSorter) Run() {
	h := algorithm.NewQuick(&MinInt64Slice{s.sortingArray})
	h.Sort()
}

func NewSingle(src []int64) Sorter {
	return &SingleSorter{sortingArray: src}
}

type ConcurrentSorter struct {
	sortedChan chan *MinInt64Slice
	sortingArray []int64
	pool *pool.Pool
	taskNum int
}

func NewConcurrent(src []int64) Sorter {
	taskNum := runtime.NumCPU()
	return &ConcurrentSorter{
		sortedChan:   make(chan *MinInt64Slice, 1),
		sortingArray: src,
		pool : nil,
		taskNum:      0,
	}
}

func (m *ConcurrentSorter) Run() {

}