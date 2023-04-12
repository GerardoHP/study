package sorting_test

import (
	"study/sorting"
	"study/utils"
	"testing"
)

func TestMergeSort_Solution(t *testing.T) {
	// arrange
	exercise := sorting.MergeSort{}
	expectedOutput := []int{1, 3, 3, 5, 6, 6}

	// act
	sln := exercise.Solution([]int{3, 6, 1, 5, 3, 6})

	// assert
	if !utils.SlicesEqual(expectedOutput, sln) {
		t.Errorf("expected %v, but found %v ", expectedOutput, sln)
		t.Fail()
	}

}
