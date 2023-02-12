package Lists

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func NewListNodeFromValue(d int) *ListNode {
	return &ListNode{Value: d}
}

func NewListNodeFromSlice(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	rln := NewListNodeFromValue(arr[0])
	for _, v := range arr[1:] {
		rln.Append(v)
	}

	return rln
}

func (ln *ListNode) Append(d int) {
	l := NewListNodeFromValue(d)
	n := ln
	for n.Next != nil {
		n = n.Next
	}

	n.Next = l
}

func (ln *ListNode) Print() {
	for ln != nil {
		fmt.Printf(" %v,", ln.Value)
		ln = ln.Next
	}

	fmt.Println()
}
