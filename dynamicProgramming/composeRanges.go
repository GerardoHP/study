package dynamicProgramming

import (
	"fmt"
	"strconv"
)

type ComposeRanges struct{}

func (ComposeRanges) Solution(nums []int) []string {
	r := []string{}
	solve([]int{}, nums, &r)
	return r
}

func solve(current, missing []int, r *[]string) {
	if len(missing) < 2 {
		if len(current) == 1 {
			*r = append(*r, strconv.Itoa(current[0]))
		}

		return
	}

	lC := len(current)
	if lC != 0 && current[lC-1]+1 != missing[0] {
		*r = append(*r, fmt.Sprint(current))
		return
	}

	m := missing[0]
	newMissing := missing[1:]
	newCurrent := append(current, m)
	solve(newCurrent, newMissing, r)
}
