package trees

import (
	"fmt"
	"study/interfaces"
)

type HasPathWithGivenSum struct {
}

var _ interfaces.Exercise = (*HasPathWithGivenSum)(nil)

func solution(t *Tree, s int) bool {
	// t.InOrderTraversal()
	// return getSum(t, s, 0)
	// t.PreOrderTraversal()
	t.LevelOrderTraversal()
	return false
}

func getSum(t *Tree, num, sum int) bool {
	if t != nil {
		if t.IsLeaf() {
			return num == (sum + t.Value.(int))
		}

		return getSum(t.Left, num, sum+t.Value.(int)) || getSum(t.Right, num, sum+t.Value.(int))
	}

	return false
}

func (t *Tree) IsLeaf() bool {
	return t.Left == nil && t.Right == nil
}

func iterateTillEnd(t *Tree) int {
	if t.Left == nil && t.Right == nil {
		fmt.Println("al fin!!!")
		return t.Value.(int)
	} else {
		return iterateTillEnd(t)
	}
}

func (HasPathWithGivenSum) Execute() {

	t := &Tree{
		Value: 4,
		Left: &Tree{
			Value: 1,
			Left: &Tree{
				Value: -2,
				Right: &Tree{
					Value: 3,
				},
			},
		},
		Right: &Tree{
			Value: 3,
			Left: &Tree{
				Value: 1,
			},
			Right: &Tree{
				Value: 2,
				Left: &Tree{
					Value: -4,
				},
				Right: &Tree{
					Value: -3,
				},
			},
		},
	}

	fmt.Printf("el resultado es %v \n", solution(t, 7))
}
