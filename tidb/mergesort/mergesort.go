package main

import (
	"fmt"
	"pingcap/talentplan/tidb/mergesort/CompBinTree"
	"pingcap/talentplan/tidb/mergesort/PQueue_Heap"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	ConcurrentSort(src)
	entrys := make([]*CompBinTree.Entry, 0)
	for i := 0; i < len(src); i++ {
		var entry = CompBinTree.NewEntry(src[i], src[i])
		entrys = append(entrys, entry)
	}
	comp := &PQueue_Heap.Comparator_default{}
	heap := PQueue_Heap.NewPQueue_Heap(comp, entrys)
	for i:= 0; !heap.IsEmpty(); i++ {
		src[i] = heap.DelMin().GetKey()
		//fmt.Println(src[i])
	}
}

func main() {
	//MergeSort(lens)
	//ConcurrentSort(lens)

	fmt.Println("start sorting")
	entrys := make([]*CompBinTree.Entry, 0)
	lens := []int64{1024, 1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1 << 13, 1 << 17, 1 << 19, 1 << 20, 1 }
	for i := 0; i < len(lens); i++ {
		var entry = CompBinTree.NewEntry(lens[i], lens[i])
		entrys = append(entrys, entry)
	}
	comp := &PQueue_Heap.Comparator_default{}
	heap := PQueue_Heap.NewPQueue_Heap(comp, entrys)
	for i:= 0; !heap.IsEmpty();i++  {
		fmt.Println(heap.DelMin().GetKey())
	}
}
