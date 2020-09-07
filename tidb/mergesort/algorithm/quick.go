package algorithm

import "sort"

type Sorter interface {
	Sort()
	GetSortable() SortableSlice
}

type Quick struct {
	slice SortableSlice
}

func NewQuick(p SortableSlice) *Quick {
	return &Quick{slice: p}
}

func (q *Quick) Sort() {
	if q.slice.Len() <= 1 {
		return
	}

	s := q.slice.GetSlice()
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func (q *Quick) GetSortable() SortableSlice {
	return q.slice
}

// 快排算法
func (q *Quick) quick_sort(left, right int) {
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right; i++ {
		if q.slice.Less(i, left) {
			explodeIndex++;
			if i != explodeIndex {
				q.slice.Swap(i, explodeIndex)
			}
		}
	}

	q.slice.Swap(left, explodeIndex)
	q.quick_sort(left, explodeIndex-1)
	q.quick_sort(explodeIndex+1, right)
}


