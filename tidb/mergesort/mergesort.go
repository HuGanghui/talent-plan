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

	initEntrySlice := make([]*CompBinTree.Entry, 0, taskNum)
	sliceArray := make([]*CompBinTree.SortedSlice, 0, taskNum)
	for i:= 0; i<taskNum; i++  {
		sortedSlice := CompBinTree.NewSortedSlice(src[intervalSlice[i]: intervalSlice[i+1]])
		sliceArray = append(sliceArray, sortedSlice)
		var entry = CompBinTree.NewEntry(sortedSlice.Value(), sortedSlice)
		initEntrySlice = append(initEntrySlice, entry)
	}

	comp := &PQueue_Heap.Comparator_default{}
	heap := PQueue_Heap.NewPQueue_Heap(comp, initEntrySlice)
	result := make([]int64, 0, len(src))
	for i:= 0; stopSignal(sliceArray, intervalSlice, taskNum); i++ {
		delMin := heap.AutoChangeRoot()
		result = append(result, delMin)
		// 打印小样例，方便测试
		//fmt.Println(result[i])
	}
	for ; !heap.IsEmpty(); {
		delMin := heap.DelMin().GetKey()
		result = append(result, delMin)
		//fmt.Println(delMin)
	}
	for i:=0; i<len(src); i++ {
		src[i] = result[i]
	}
 }

 func stopSignal(sliceArray []*CompBinTree.SortedSlice, intervalSlice []int, taskNum int) bool {
 	var result = false
	for i:= 0; i < taskNum; i++{
		result = result || sliceArray[i].HasNext()
	}
	return result
 }

func main() {
	// 小样例，方便测试
	fmt.Println("start sorting")
	lens := []int64{1024, 1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1 << 13, 1 << 17, 1 << 19, 1 << 20, 1, 1 << 21 }
	MergeSort(lens)
}
