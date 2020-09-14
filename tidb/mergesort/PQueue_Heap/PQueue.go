package PQueue_Heap

import "pingcap/talentplan/tidb/mergesort/CompBinTree"

type PQueue interface {
	getSize() int64
	isEmpty() bool
	getMin() *CompBinTree.Entry
	insert(key int64, value int64) *CompBinTree.Entry
	delMin() *CompBinTree.Entry
}
