package algorithm

type SortableSlice interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Append(v int64)
	IndexOf(int) (int64, error)
	Pop() (int64, error)
	GetSlice() []int64
}


