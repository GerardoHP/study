package trees

import (
	"fmt"
	"strconv"
	"study/interfaces"
)

type DigitTreeSum struct {
}

func (d DigitTreeSum) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*DigitTreeSum)(nil)

func (DigitTreeSum) Solution(t *Tree) int {
	var nums []string
	sln(t, "", &nums)
	return sum(nums)
}

func sln(t *Tree, str string, nums *[]string) {
	if t != nil {
		nStr := fmt.Sprintf("%v%v", str, t.Value)
		sln(t.Left, nStr, nums)

		if t.IsLeaf() {
			*nums = append(*nums, nStr)
		}

		sln(t.Right, nStr, nums)
	}
}

func sum(nums []string) int {
	sum := 0
	for _, str := range nums {
		n, _ := strconv.Atoi(str)
		sum += n
	}

	return int(sum)
}
