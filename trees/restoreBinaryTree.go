package trees

import (
	"fmt"
	"study/interfaces"
)

type RestoreBinaryTree struct {
}

var _ interfaces.Exercise = (*RestoreBinaryTree)(nil)

func (RestoreBinaryTree) solution(inorder []int, preorder []int) *Tree {
	return solutionRestoreBinaryTree(inorder, preorder)
}

func solutionRestoreBinaryTree(inorder []int, preorder []int) *Tree {
	if len(inorder) == 0 {
		return nil
	}

	i := getMiddlePosition(preorder[0], inorder)
	leftInOrder := inorder[:i]
	rightInOrder := inorder[i+1:]
	l := len(leftInOrder)
	leftPreOrder := preorder[1 : l+1]
	rightPreOrder := preorder[l+1:]

	tree := &Tree{
		Value: preorder[0],
		Left:  solutionRestoreBinaryTree(leftInOrder, leftPreOrder),
		Right: solutionRestoreBinaryTree(rightInOrder, rightPreOrder),
	}

	return tree
}

func getMiddlePosition(search int, inorder []int) int {
	for i, v := range inorder {
		if v == search {
			return i
		}
	}

	return -1
}

func (r RestoreBinaryTree) Execute() {
	inorder := []int{4, 2, 1, 5, 3, 6}
	preorder := []int{1, 2, 4, 3, 5, 6}
	//
	//inorder := []int{-100000, -99999, -99998}
	//preorder := []int{-99999, -100000, -99998}
	result := r.solution(inorder, preorder)
	fmt.Printf("the solution is %v \n", result)
}
