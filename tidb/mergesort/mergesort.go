package main

import (
	"fmt"
	"pingcap/talentplan/tidb/mergesort/CompBinTree"
	"pingcap/talentplan/tidb/mergesort/PQueue_Heap"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	intervalSlice, taskNum:= ConcurrentSort(src)
	// 当src元素小于默认阈值，就使用内置的快排算法
	if intervalSlice == nil {
		return
	}

	pointSlice := make([]int, 0, taskNum)
	initEntrySlice := make([]*CompBinTree.Entry, 0, taskNum)
	for i:= 0; i<taskNum; i++  {
		pointSlice = append(pointSlice, intervalSlice[i])
		var entry = CompBinTree.NewEntry(src[pointSlice[i]], int64(i))
		initEntrySlice = append(initEntrySlice, entry)
	}

	comp := &PQueue_Heap.Comparator_default{}
	heap := PQueue_Heap.NewPQueue_Heap(comp, initEntrySlice)
	result := make([]int64, 0, len(src))
	for i:= 0; !stopSignal(pointSlice, intervalSlice, taskNum); i++ {
		//delMin := heap.DelMin()
		delMin := heap.GetRoot()
		result = append(result, delMin.GetKey())
		insert(heap, src, pointSlice, intervalSlice, delMin)
		// 打印小样例，方便测试
		//fmt.Println(result[i])
	}
	for i:=0; i<len(src); i++ {
		src[i] = result[i]
	}
 }

 func stopSignal(pointSlice []int, intervalSlice []int, taskNum int) bool {
 	var result = true
	for i:= 0; i < taskNum; i++{
		result = result && pointSlice[i] >= intervalSlice[i+1]
	}
	return result
 }

func insert(heap *PQueue_Heap.PQueue_Heap, src []int64,
	pointSlice []int, intervalSlice []int, entry *CompBinTree.Entry)  {
	index := entry.GetValue()
	pointSlice[index]++
	if pointSlice[index] < intervalSlice[index+1] {
		var entry = CompBinTree.NewEntry(src[pointSlice[index]], index)
		heap.ReplaceRoot(entry)
	} else {
		heap.DelMin()
	}
}

func main() {
	// 小样例，方便测试
	fmt.Println("start sorting")
	lens := []int64{1024, 1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1 << 13, 1 << 17, 1 << 19, 1 << 20, 1, 1 << 21 }
	MergeSort(lens)
}
