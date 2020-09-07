package task

import (
	"log"
	"pingcap/talentplan/tidb/mergesort/algorithm"
)

type SortTask struct {
	sorter  *algorithm.Quick
	retChan chan *MinInt64Slice
}

func NewSortTask(src []int64, retChan chan *MinInt64Slice) *SortTask {

	return &SortTask{
		sorter: algorithm.NewQuick(&MinInt64Slice{
			array: src,
		}),
		retChan: retChan,
	}
}

func (s *SortTask) Run() error {
	s.sorter.Sort()
	s.retChan <- s.sorter.GetSortable().(*MinInt64Slice)
	return nil
}

func (s *MinInt64Slice) RunError(err error) {
	log.Println("sort task run error", err)
}

type MergeTask struct {
	slices  [][]int64
	retChan chan []int64
}

func NewMergeTask(slices [][]int64, retChan chan []int64) *MergeTask {
	return &MergeTask{
		slices:  slices,
		retChan: retChan,
	}
}

func (m *MergeTask) Run() error {
	sortedSlices := make([]*algorithm.SortableSlice, 0, len(m.slices))
	for _, s := range m.slices {
		sortedSlices = append(sortedSlices, algorithm.NewSortedSlice(s))
	}
	merge := algorithm.NewHeapMerge(sortedSlices)
	m.retChan <- merge.Sort()
	return nil
}

func (m *MergeTask) RunError(err error) {
	log.Println("MergeTask run error", err)
}
