package backtracking_test

import (
	"study/backtracking"
	"study/utils"
	"testing"
)

func TestSumSubsets_Solution(t *testing.T) {
	// arrange
	arr := []int{1, 2, 3, 4, 5}
	num := 5
	r := [][]int{{1, 4},
		{2, 3},
		{5}}
	exercise := backtracking.SumSubsets{}

	// act
	sln := exercise.Solution(arr, num)

	// assert
	for k := range sln {
		a := sln[k]
		b := r[k]
		if !utils.SlicesEqual(a, b) {
			t.Errorf("expected %v, but found %v in test 1", b, a)
			t.Fail()
		}
	}
}
