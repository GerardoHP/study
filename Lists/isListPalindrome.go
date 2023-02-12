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
	for slow != nil {
		fmt.Printf("slow: %v \n", slow.Value)
		fast := slow
		for fast != nil || (fast.Next != nil && fast.Next.Next != nil) {
			fmt.Printf("fast : %v \n", fast.Value)
			if fast.Next == nil {
				fmt.Println("vamonos")
				if slow.Value != fast.Value {
					return false
				}

				fast = nil
				break
			}

			if fast.Next.Next == nil {
				fast = fast.Next
			} else {
				fast = fast.Next.Next
			}
		}

		slow = slow.Next
	}

	return true
}

func (IsListPalindrome) Execute() {
	//arr := []int{0, 1, 0}
	arr := []int{1, 2, 2, 3}
	l := NewListNodeFromSlice(arr)
	fmt.Printf("el resultado es %v \n", solution(l))
	//l.Print()
}
