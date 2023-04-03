package backtracking

import (
	"strconv"
	"strings"
	"study/interfaces"
)

type IsPasswordCorrect struct{}

func (i IsPasswordCorrect) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*IsPasswordCorrect)(nil)

func (IsPasswordCorrect) Solution(pass string, nums []int) bool {
	return solveIsPasswordCorrect(pass, nums, []int{})
}

func solveIsPasswordCorrect(pass string, nums, current []int) bool {
	if len(nums) == 0 {
		str := intSliceToString(current)
		if pass == str {
			return true
		}
	}

	for k, v := range nums {
		newCurrent := append(current, v)

		newNums := make([]int, len(nums)-1)
		copy(newNums[:k], nums[:k])
		copy(newNums[k:], nums[k+1:])
		b := solveIsPasswordCorrect(pass, newNums, newCurrent)
		if b {
			return true
		}
	}

	return false
}

func intSliceToString(nums []int) string {
	strBuilder := strings.Builder{}
	for _, v := range nums {
		strBuilder.WriteString(strconv.Itoa(v))
	}

	return strBuilder.String()
}
