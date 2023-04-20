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

	solve([]int{}, nums, &r)
	return r
}

func solve(current, missing []int, r *[]string) {
	lC := len(current)
	if lC != 0 && current[lC-1]+1 != missing[0] {
		*r = append(*r, sliceToRange(current))
		solve([]int{}, missing, r)
		return
	}

	if len(missing) == 1 {
		*r = append(*r, strconv.Itoa(missing[0]))
		return
	}

	m := missing[0]
	newMissing := missing[1:]
	newCurrent := append(current, m)
	solve(newCurrent, newMissing, r)
}

func sliceToRange(s []int) string {
	l := len(s)
	return fmt.Sprintf("%v->%v", s[0], s[l-1])
}
