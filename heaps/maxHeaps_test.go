package heaps_test

import (
	"study/heaps"
	"testing"
)

type entryHeaps struct {
	input          []int
	expectedOutput []int
}

func TestMaxHeap_Insert(t *testing.T) {
	entries := []entryHeaps{
		{input: []int{10, 20, 30}, expectedOutput: []int{30, 20, 10}},
	}

	for _, entr := range entries {
		r := heaps.NewMaxHeap()
		for _, v := range entr.input {
			r.Insert(v)
		}

		for _, v := range entr.expectedOutput {
			if v != r.Extract() {
				t.Fail()
			}
		}
	}
}
