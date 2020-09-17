package PQueue_Heap

import (
	"pingcap/talentplan/tidb/mergesort/CompBinTree"
)

type PQueue interface {
	GetSize() int64
	IsEmpty() bool
	GetMin() *CompBinTree.Entry
	Insert(entry *CompBinTree.Entry) *CompBinTree.Entry
	DelMin() *CompBinTree.Entry
}

type PQueue_Heap struct {
	heap *CompBinTree.CompBinTree_Vector
	comp Comparator
}

func NewPQueue_Heap(c Comparator, entrys []*CompBinTree.Entry) *PQueue_Heap {
	heap := CompBinTree.NewCompBinTree_Vector(entrys)
	pqueue := &PQueue_Heap{
		heap: heap,
		comp: c,
	}
	for i := pqueue.heap.GetSize() / 2 - 1 ; i >= 0; i--  {
		pqueue.percolateDown(pqueue.heap.PosOfNode(i))
	}
	return pqueue
}

func (heap *PQueue_Heap) GetSize() int64 {
	return heap.heap.GetSize()
}

func (heap *PQueue_Heap) IsEmpty() bool {
	return heap.heap.IsEmpty()
}

func (heap *PQueue_Heap) GetMin() *CompBinTree.Entry {
	return heap.heap.GetRoot().GetElem()
}

func (heap *PQueue_Heap) Insert(entry *CompBinTree.Entry) *CompBinTree.Entry {
	heap.percolateUp(heap.heap.AddLast(entry))
	return entry
}

func (heap *PQueue_Heap) DelMin() *CompBinTree.Entry{
	if heap.IsEmpty() {
		return nil
	}
	min := heap.heap.GetRoot().GetElem()
	if 1 == heap.heap.GetSize() {
		heap.heap.DelLast()
	} else {
		heap.heap.GetRoot().SetElem(heap.heap.DelLast().GetElem())
		heap.percolateDown(heap.heap.GetRoot())
	}
	return min
}

func (heap *PQueue_Heap) GetRoot() *CompBinTree.Entry {
	if heap.IsEmpty() {
		return nil
	}
	return heap.heap.GetRoot().GetElem()
}

func (heap *PQueue_Heap) AutoChangeRoot() int64 {
	ele := heap.heap.GetRoot().GetElem()
	rootValue := ele.GetKey()
	sortedSlice := ele.GetValue()
	if sortedSlice.HasNext() {
		sortedSlice.Next()
		ele.SetKey(sortedSlice.Value())
		heap.percolateDown(heap.heap.GetRoot())
	} else {
		heap.DelMin()
	}
	return rootValue
}

func (heap *PQueue_Heap) ReplaceRoot(entry *CompBinTree.Entry) {
	heap.heap.GetRoot().SetElem(entry)
	heap.percolateDown(heap.heap.GetRoot())
}

func (heap *PQueue_Heap) swapParentChild(u CompBinTree.BinTreePosition,
	v CompBinTree.BinTreePosition) {
	temp := u.GetElem()
	u.SetElem(v.GetElem())
	v.SetElem(temp)
}

func (heap *PQueue_Heap) percolateUp(v CompBinTree.BinTreePosition) {
	root := heap.heap.GetRoot()
	for ; v != root; {
		p := v.GetParent()
		if 0 > heap.comp.Compare(p.GetElem(), v.GetElem()) {
			break
		}
		heap.swapParentChild(p, v)
		v = p
	}
}

func (heap *PQueue_Heap) percolateDown(v CompBinTree.BinTreePosition) {
	for ; v.HasLChild(); {
		smallerChild := v.GetLChild()
		if v.HasRChild() && 0 < heap.comp.Compare(v.GetLChild().GetElem(),
			v.GetRChild().GetElem()) {
			smallerChild = v.GetRChild()
		}
		if 0 < heap.comp.Compare(smallerChild.GetElem(), v.GetElem()) {
			break
		}
		heap.swapParentChild(v, smallerChild)
		v = smallerChild
	}
}
