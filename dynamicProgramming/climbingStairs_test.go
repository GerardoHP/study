package dynamicProgramming_test

import (
	"study/dynamicProgramming"
	"testing"
)

func TestExecute(t *testing.T) {
	// arrange
	tests := map[int]int{
		5:  8,
		10: 89,
	}

	// act, assert
	exercise := dynamicProgramming.ClimbingStairs{}
	for val, expected := range tests {
		sln := exercise.Solution(val)
		if expected != sln {
			t.Errorf("expected %v, but found %v \n", expected, sln)
			t.Fail()
		}
	}

}
