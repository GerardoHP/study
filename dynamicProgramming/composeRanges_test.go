package dynamicProgramming_test

import (
	"study/dynamicProgramming"
	"study/utils"
	"testing"
)

func TestComposeRanges_Solution_1(t *testing.T) {
	// arrange
	input := []int{-1, 0, 1, 2, 6, 7, 9}
	output := []string{"-1->2", "6->7", "9"}

	// act, assert
	testComposeRangesSolution(t, input, output, "test 1")
}

func TestComposeRanges_Solution_6(t *testing.T) {
	// arrange
	input := []int{0, 1}
	output := []string{"0->1"}

	// act, assert
	testComposeRangesSolution(t, input, output, "test 1")
}

func testComposeRangesSolution(t *testing.T, input []int, output []string, test string) {
	// act
	exercise := dynamicProgramming.ComposeRanges{}
	sln := exercise.Solution(input)

	// assert
	if !utils.StringSlicesEqual(output, sln) {
		t.Errorf("expected %v, but found %v in %v \n", output, sln, test)
		t.Fail()
	}
}
