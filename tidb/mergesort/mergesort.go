package main

import (
	"fmt"
	"pingcap/talentplan/tidb/mergesort/CompBinTree"
	"pingcap/talentplan/tidb/mergesort/PQueue_Heap"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	var intervalSlice = ConcurrentSort(src)
	// 当src元素小于4个，就使用内置的快排算法
	if intervalSlice == nil {
		return
	}

	pointSlice := make([]int, 0, 4)
	initEntrySlice := make([]*CompBinTree.Entry, 0, 4)
	for i:= 0; i<4; i++  {
		pointSlice = append(pointSlice, intervalSlice[i])
		var entry = CompBinTree.NewEntry(src[pointSlice[i]], int64(i))
		initEntrySlice = append(initEntrySlice, entry)
	}

	comp := &PQueue_Heap.Comparator_default{}
	heap := PQueue_Heap.NewPQueue_Heap(comp, initEntrySlice)
	result := make([]int64, 0, len(src))
	for i:= 0; !heap.IsEmpty(); i++ {
		delMin := heap.DelMin()
		result = append(result, delMin.GetKey())
		insert(heap, src, pointSlice, intervalSlice, delMin)
		// 打印小样例，方便测试
		//fmt.Println(result[i])
	}
	for i:=0; i<len(src); i++ {
		src[i] = result[i]
	}
 }

func insert(heap PQueue_Heap.PQueue, src []int64,
	pointSlice []int, intervalSlice []int, entry *CompBinTree.Entry)  {
	index := entry.GetValue()
	pointSlice[index]++
	if pointSlice[index] < intervalSlice[index+1] {
		var entry = CompBinTree.NewEntry(src[pointSlice[index]], index)
		heap.Insert(entry)
	}
}

func main() {
	// 小样例，方便测试
	fmt.Println("start sorting")
	lens := []int64{1024, 1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1 << 13, 1 << 17, 1 << 19, 1 << 20, 1 }
	MergeSort(lens)
}
