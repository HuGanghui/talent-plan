package PQueue_Heap

import "pingcap/talentplan/tidb/mergesort/CompBinTree"

type Comparator interface {
	compare(a CompBinTree.Entry, b CompBinTree.Entry) int64
}
