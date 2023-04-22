package dynamicProgramming

import (
	"fmt"
	"strconv"
)

type ComposeRanges struct{}

func (ComposeRanges) Solution(nums []int) []string {
	r := []string{}

	r = solve(nums)
	return r
}

func solve(nums []int) []string {
	r := []string{}
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			i++
		}

		end := nums[i]
		rangeStr := strconv.Itoa(start)
		if start != end {
			rangeStr = fmt.Sprintf("%v->%v", start, end)
		}

		r = append(r, rangeStr)
	}

	return r
}
