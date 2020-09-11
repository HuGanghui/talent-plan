package main

import (
	"fmt"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	ConcurrentSort(src)
}

func main() {
	lens := []int64{1024, 1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1 << 13, 1 << 17, 1 << 19, 1 << 20, 1 }
	//MergeSort(lens)
	ConcurrentSort(lens)
	for _, ele := range lens {
		fmt.Println(ele)
	}
}
