package trees

import (
	"study/interfaces"
	"study/utils"
)

type LargestValuesInTreeRows struct {
}

func (l LargestValuesInTreeRows) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*LargestValuesInTreeRows)(nil)

func (LargestValuesInTreeRows) Solution(t *Tree) []int {
	var a []int
	breadthLevel(&a, t, 0)
	return a
}

func breadthLevel(res *[]int, t *Tree, level int) {
	if t == nil {
		return
	}

	val := int(t.Value.(float64))
	if len(*res) == level {
		*res = append(*res, val)
	} else {
		(*res)[level] = utils.MaxInt((*res)[level], val)
	}

	breadthLevel(res, t.Left, level+1)
	breadthLevel(res, t.Right, level+1)
}
