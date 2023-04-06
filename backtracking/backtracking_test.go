package backtracking_test

import (
	"study/backtracking"
	"testing"
)

func TestBacktracking_Solution(t *testing.T) {
	// arrange
	expectedCount := 6
	nums := []int{1, 2, 3}

	// act
	exercise := backtracking.Backtracking{}
	sln := exercise.Solution(nums)

	// assert
	if len(sln) != expectedCount {
		t.Errorf("expected %v but found %v \n", expectedCount, len(sln))
		t.Fail()
	}
}
