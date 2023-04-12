package backtracking

import (
	"fmt"
	"sort"
	"study/interfaces"
)

type CombinationSum struct {
}

func (c CombinationSum) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*CombinationSum)(nil)

func (CombinationSum) Solution(a []int, sum int) string {
	sums := [][]int{}
	str := ""
	sort.Ints(a)
	aAux := []int{}
	m := make(map[int]bool)
	for _, v := range a {
		if _, found := m[v]; !found {
			m[v] = true
			aAux = append(aAux, v)
		}
	}

	solveCombinationSum(aAux, []int{}, sum, &sums)
	for _, v1 := range sums {
		str = fmt.Sprintf("%v%v", str, "(")
		for y, v2 := range v1 {
			str = fmt.Sprintf("%v%v", str, v2)
			if y < len(v1)-1 {
				str = fmt.Sprintf("%v ", str)
			}
		}
		str = fmt.Sprintf("%v%v", str, ")")
	}
	return str
}

func solveCombinationSum(a, currentSum []int, sum int, sums *[][]int) {
	if sum == 0 {
		*sums = append(*sums, currentSum)
		return
	}

	for _, v := range a {
		latest := 0
		if l := len(currentSum); l > 0 {
			latest = currentSum[l-1]
		}

		if v <= sum && v >= latest {
			newSum := make([]int, len(currentSum))
			copy(newSum, currentSum)
			newSum = append(newSum, v)
			solveCombinationSum(a, newSum, sum-v, sums)
		}
	}

}
