package sorting

import "study/interfaces"

type MergeSort struct {
}

func (m MergeSort) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*MergeSort)(nil)

func merge(sequence []int, left int, middle int, right int) {
	result := make([]int, right-left, right-left)
	i := left
	j := middle
	p := 0
	for i < middle && j < right {
		if sequence[i] < sequence[j] {
			result[p] = sequence[i]
			p++
			i++
		} else {
			result[p] = sequence[j]
			p++
			j++
		}
	}
	for i < middle {
		result[p] = sequence[i]
		p++
		i++
	}
	for j < right {
		result[p] = sequence[j]
		p++
		j++
	}
	for i := left; i < right; i++ {
		sequence[i] = result[i-left]
	}
}

func split(sequence []int, left int, right int) {
	// ...
	if right-left <= 1 {
		return
	}

	middle := (left + right) / 2
	split(sequence, left, middle)
	split(sequence, middle, right)
	merge(sequence, left, middle, right)
}

func (MergeSort) Solution(sequence []int) []int {
	split(sequence, 0, len(sequence))
	return sequence
}
