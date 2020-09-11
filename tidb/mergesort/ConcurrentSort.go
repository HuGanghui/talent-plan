package main

import (
	"runtime"
	"sort"
)

func ConcurrentSort(src []int64)  {
	taskNum := runtime.NumCPU()
	start := 0
	count := 1
	array_length := len(src)
	step := array_length / taskNum
	chs := make([]chan int, taskNum)
	interval := make([]int, 0)
	interval = append(interval, start)
	for ; start < array_length; {
		end := start + step
		if taskNum == count {
			end = array_length
		}
		interval = append(interval, end)
		chs[count-1] = make(chan int)
		go Sort(src[start:end], chs[count-1])
		start = end
		count++
	}
	for _, ch := range chs {
		<-ch
	}
	Sort2(src)
}

func Sort(src []int64, succ chan int) {
	sort.Slice(src, func(i, j int) bool {
		return src[i] < src[j]
	})
	succ <- 1
}

func Sort2(src []int64) {
	sort.Slice(src, func(i, j int) bool {
		return src[i] < src[j]
	})
}