package queue

import (
	"study/interfaces"
	"study/stack"
)

type NearestGreater struct {
}

func (n NearestGreater) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*NearestGreater)(nil)

func (NearestGreater) Solution(a []int) []int {
	var b []int
	for i, val := range a {
		left := stack.NewStackFromSlice(a[:i])
		right := NewQueueFromSlice(a[i+1:])
		v := GetLower(*left, *right, val, i)
		b = append(b, v)
	}

	return b
}

func GetLower(left stack.Stack, right Queue, val, index int) int {
	max := 0
	for {
		if left.IsEmpty() && right.IsEmpty() {
			break
		}

		max++
		if left.IsNotEmpty() {
			if x := left.Pop(); x.(int) > val {
				return index - max
			}
		}

		if right.IsNotEmpty() {
			if x := right.Dequeue(); x.(int) > val {
				return max + index
			}
		}
	}

	return -1
}
