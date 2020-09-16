package PQueue_Heap

import "pingcap/talentplan/tidb/mergesort/CompBinTree"

type Comparator interface {
	Compare(a *CompBinTree.Entry, b *CompBinTree.Entry) int
}

type Comparator_default struct {}

func (comp Comparator_default) Compare(a *CompBinTree.Entry, b *CompBinTree.Entry) int {
	if a.GetKey() >= b.GetKey() {
		return 1
	} else {
		return -1
	}
}
