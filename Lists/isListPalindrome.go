package Lists

import (
	"fmt"
	"study/Interfaces"
)

type IsListPalindrome struct {
}

var _ Interfaces.Exercise = (*IsListPalindrome)(nil)

func solution(l *ListNode) bool {
	slow := l
	fast := l
	i := 0
	for fast.Next != nil {
		fast = fast.Next.Next
		if fast == nil {
			break
		}

		slow = slow.Next
		i++
	}

	slowR := reverseList(slow.Next)
	for l != nil && slowR != nil {
		if l.Value != slowR.Value {
			return false
		}

		l = l.Next
		slowR = slowR.Next
	}

	return true
}

func reverseList(n *ListNode) *ListNode {
	var prev *ListNode = nil
	current := n
	var next *ListNode = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func (IsListPalindrome) Execute() {
	// arr := []int{0, 1, 0}
	arr := []int{1, 2, 2, 3}
	// arr := []int{1, 1000000000, -1000000000, -1000000000, 1000000000, 1}
	l := NewListNodeFromSlice(arr)
	// l, _ := NewListFromFile()
	fmt.Printf("el resultado es %v \n", solution(l))
	//l.Print()
}
