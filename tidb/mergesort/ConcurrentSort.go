package main

import (
	"runtime"
	"sort"
)

const SortingArrayThreshold = 9

func ConcurrentSort(src []int64) (interval []int, taskNum int) {
	taskNum = runtime.NumCPU()
	start := 0
	count := 1
	arrayLength := len(src)
	interval = make([]int, 0, taskNum+1)
	if arrayLength < SortingArrayThreshold {
		SortDefault(src)
		interval = nil
	} else {
		step := arrayLength / taskNum
		chs := make([]chan int, taskNum)
		interval = append(interval, start)
		for ; start < arrayLength; {
			end := start + step
			if taskNum == count {
				end = arrayLength
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
	}
	return interval, taskNum
}

func Sort(src []int64, succ chan int) {
	sort.Slice(src, func(i, j int) bool {
		return src[i] < src[j]
	})
	succ <- 1
}

func SortDefault(src []int64) {
	sort.Slice(src, func(i, j int) bool {
		return src[i] < src[j]
	})
}