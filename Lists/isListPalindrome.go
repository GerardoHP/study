package Lists

import (
	"study/Interfaces"
)

type IsListPalindrome struct {
}

var _ Interfaces.Exercise = (*IsListPalindrome)(nil)

func solution(l *ListNode) bool {
	slow := l
	for slow != nil {
		// fmt.Printf("slow: %v \n", slow.Value)
		fast := slow
		if fast == nil {
			return true
		}

		for fast.Next != nil {
			// fmt.Printf("fast: %v \n", fast.Value)

			if fast.Next.Next == nil {
				// este el el ultimo para comprara
				if fast.Next.Value != slow.Value {
					return false
				}

				// se elimina
				fast.Next = nil
				break
			}

			if fast.Next.Next.Next == nil {
				if fast.Next.Next.Value != slow.Value {
					return false
				}

				fast.Next.Next = nil
				break
			}

			fast = fast.Next.Next
		}

		slow = slow.Next
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
	// arr := []int{0, 1, 2, 1, 0}
	arr := []int{1, 2, 2, 3}
	// arr := []int{1, 1000000000, -1000000000, -1000000000, 1000000000, 1}
	l := NewListNodeFromSlice(arr)
	l2 := reverseList(l)
	l2.Print()
	// l, _ := NewListFromFile()
	// fmt.Printf("el resultado es %v \n", solution(l))
	//l.Print()
}
