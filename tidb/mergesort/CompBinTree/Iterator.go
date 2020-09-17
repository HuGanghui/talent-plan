package CompBinTree

type Iterator struct {
	slice []int64
	index int
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.slice)-1
}

func (i *Iterator) Next() {
	i.index++
}

func (i *Iterator) Value() int64 {
	return i.slice[i.index]
}

type SortedSlice struct {
	slice []int64
	Iterator
}

func NewSortedSlice(slice []int64) *SortedSlice {
	return &SortedSlice{
		slice: slice,
		Iterator: Iterator{
			slice: slice,
			index: 0,
		},
	}
}
