package heaps

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap []int

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func NewMaxHeapFromSlice(slc []int) *MaxHeap {
	mh := &MaxHeap{}
	for _, s := range slc {
		mh.Insert(s)
	}

	return mh
}

// Insert adds on element to the heap
func (h *MaxHeap) Insert(key int) {
	*h = append(*h, key)
	h.maxHeapifyUp(len(*h) - 1)
}

// Extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := (*h)[0]
	l := len((*h)) - 1
	if l == -1 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	(*h)[0] = (*h)[l]
	(*h) = (*h)[:l]
	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for p := parent(index); (*h)[p] < (*h)[index]; {
		h.swap(p, index)
		index = p
		p = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len((*h)) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		switch {
		case l == lastIndex:
			childToCompare = l
		case (*h)[l] > (*h)[r]:
			childToCompare = l
		default:
			childToCompare = r
		}

		if (*h)[index] < (*h)[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	(*h)[i1], (*h)[i2] = (*h)[i2], (*h)[i1]
}

func parent(i int) int {
	if i == 0 {
		return 0
	}

	p := (i - 1) / 2
	return p
}

func left(p int) int {
	left := p*2 + 1
	return left
}

func right(p int) int {
	right := left(p) + 1
	return right
}
