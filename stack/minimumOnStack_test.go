package stack_test

import (
	"study/stack"
	"study/utils"
	"testing"
)

type MinimumOnStackEntry struct {
	Test   string
	Input  []string
	Output []int
}

func TestMinimumOnStack_Solution(t *testing.T) {
	test := NewMinimumOnStackEntry(
		"Test 1",
		[]string{"push 10", "min", "push 5", "min", "push 8", "min", "pop", "min", "pop", "min"},
		[]int{10, 5, 5, 5, 10},
	)
	testMinimumOnStack_Solution(t, test)
}

func testMinimumOnStack_Solution(t *testing.T, test *MinimumOnStackEntry) {
	// arrange
	exercise := stack.MinimumOnStack{}

	// act
	sln := exercise.Solution(test.Input)

	// assert
	if !utils.SlicesEqual(sln, test.Output) {
		t.Errorf("expected %v, but found %v on %v \n", test.Output, sln, test.Test)
	}
}

func NewMinimumOnStackEntry(test string, input []string, output []int) *MinimumOnStackEntry {
	return &MinimumOnStackEntry{
		Test:   test,
		Input:  input,
		Output: output,
	}
}
