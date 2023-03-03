package Trees

import (
	"fmt"

	"study/Interfaces"
)

type IsTreeSymmetric struct {
}

var _ Interfaces.Exercise = (*IsTreeSymmetric)(nil)

func (IsTreeSymmetric) solution(t *Tree) bool {
	return compare(t.Left, t.Right)
}

func compare(left *Tree, right *Tree) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}

	fmt.Printf("right %v, left %v \n", left.Value, right.Value)
	if left.Value != right.Value {
		return false
	}

	return compare(left.Right, right.Left) && compare(left.Left, right.Right)
}

func (IsTreeSymmetric) Execute() {

	//t := &Tree{
	//	Value: 4,
	//	Left: &Tree{
	//		Value: 1,
	//		Left: &Tree{
	//			Value: -2,
	//			Right: &Tree{
	//				Value: 3,
	//			},
	//		},
	//	},
	//	Right: &Tree{
	//		Value: 1,
	//		Left: &Tree{
	//			Value: 3,
	//		},
	//		Right: &Tree{
	//			Value: -2,
	//			Left: &Tree{
	//				Value: -4,
	//			},
	//			//Right: &Tree{
	//			//	Value: -3,
	//			//},
	//		},
	//	},
	//}

	t := &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 1,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 3,
				},
			},
		},
		Right: &Tree{
			Value: 2,
			Right: &Tree{
				Value: 3,
				Right: &Tree{
					Value: 1,
					Right: &Tree{
						Value: 2,
					},
				},
			},
		},
	}

	e := IsTreeSymmetric{}
	fmt.Printf("el resultado es %v \n", e.solution(t))
}
