package Trees

import "study/Interfaces"

type DeleteFromBST struct{}

var _ Interfaces.Exercise = (*DeleteFromBST)(nil)

func (DeleteFromBST) Solution(t *Tree, queries []int) *Tree {
	for _, v := range queries {
		t = findAndRemove(t, v)
	}

	return t
}

func farthestRightValue(t *Tree) *Tree {
	if t.Right == nil || t.isLeaf() {
		return t
	}

	return farthestRightValue(t.Right)
}

func findAndRemove(t *Tree, find int) *Tree {
	f := float64(find)
	if t == nil {
		return t
	}

	if t.Value.(float64) > f {
		t.Left = findAndRemove(t.Left, find)
	}

	if t.Value.(float64) < f {
		t.Right = findAndRemove(t.Right, find)
	}

	if t.Value.(float64) == f {
		if t.isLeaf() {
			t = nil
			return t
		}

		if t.Left != nil {
			farthestRight := farthestRightValue(t.Left)
			t.Value = farthestRight.Value
			t.Left = findAndRemove(t.Left, int(farthestRight.Value.(float64)))
		} else {
			temp := t.Right
			t = nil
			t = temp
			return t
		}
	}

	return t
}

func (d DeleteFromBST) Execute() {
	//TODO implement me
	panic("implement me")
}
