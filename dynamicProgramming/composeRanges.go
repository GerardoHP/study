package dynamicProgramming

import (
	"fmt"
	"strconv"
)

type ComposeRanges struct{}

func (ComposeRanges) Solution(nums []int) []string {
	r := []string{}
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	solve(nums, &r)
	return r
}

func solve(nums []int, r *[]string) {
	if len(nums) == 1 {
		*r = append(*r, strconv.Itoa(nums[0]))
		return
	}

	c := 0
	x := []int{}
	for k, v := range nums {
		x = append(x, v)
		if l := len(nums); k+1 < l && v+1 != nums[k+1] {
			c = k
			break
		}
	}

	*r = append(*r, sliceToRange(x))
	if c != 0 {
		solve(nums[c+1:], r)
	}
}

func sliceToRange(s []int) string {
	l := len(s)
	return fmt.Sprintf("%v->%v", s[0], s[l-1])
}
