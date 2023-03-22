package heaps

import "study/interfaces"

type KthLargestElement struct {
}

var _ interfaces.Exercise = (*KthLargestElement)(nil)

func (KthLargestElement) Solution(nums []int, k int) int {
	maxHeap := NewMaxHeapFromSlice(nums)
	r := -1
	for i := 0; i < len(nums); i++ {
		r = maxHeap.Extract()
		if i == k-1 {
			return r
		}
	}

	return r
}

func (k KthLargestElement) Execute() {
	//TODO implement me
	panic("implement me")
}
