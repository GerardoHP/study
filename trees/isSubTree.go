package trees

import (
	"fmt"
	"study/interfaces"
)

//
// Binary trees are already defined with this interface:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }

type IsSubTree struct {
}

var _ interfaces.Exercise = (*IsSubTree)(nil)

func (IsSubTree) solution(t1 *Tree, t2 *Tree) bool {
	if t1 == t2 {
		return true
	}

	if t1 == nil {
		return false
	}

	if t2 == nil {
		return true
	}

	return searchTroughTrees(t1, t2)
}

func searchTroughTrees(t *Tree, subTree *Tree) bool {
	if t != nil {
		if compareTree(t, subTree) {
			return true
		}

		return searchTroughTrees(t.Left, subTree) || searchTroughTrees(t.Right, subTree)
	}

	return false
}

func compareTree(t1 *Tree, t2 *Tree) bool {
	// fmt.Printf(" %v %v \n", t1, t2)
	if t1 == nil || t2 == nil {
		return t1 == t2
	}

	// fmt.Printf(" %v | %v \n", t1.Value, t2.Value)
	if t1.Value != t2.Value {
		return false
	}

	return compareTree(t1.Left, t2.Left) && compareTree(t1.Right, t2.Right)
}

func (ist IsSubTree) Execute() {
	x, y, _ := NewTreesFromFile()
	fmt.Printf("The soultion is %v \n", ist.solution(x, y))
}
