package backtracking_test

import (
	"study/backtracking"
	"testing"
)

type CombinationSumEntry struct {
	Name   string
	Input  CombinationSumInput `json:"input"`
	Output string              `json:"output"`
}

type CombinationSumInput struct {
	A   []int `json:"a"`
	Sum int   `json:"sum"`
}

func NewCombinationSumEntry(name string, a []int, sum int, output string) *CombinationSumEntry {
	return &CombinationSumEntry{
		Name: name,
		Input: CombinationSumInput{
			A:   a,
			Sum: sum,
		},
		Output: output,
	}
}

func TestCombinationSum_Solution_1(t *testing.T) {
	// arrange
	entry := NewCombinationSumEntry(
		"Test 1",
		[]int{2, 3, 5, 9},
		9,
		"(2 2 2 3)(2 2 5)(3 3 3)(9)")

	// act, assert
	testCombinationSumSolution(t, *entry)
}

func TestCombinationSum_Solution_3(t *testing.T) {
	// arrange
	entry := NewCombinationSumEntry(
		"Test 3",
		[]int{8, 1, 8, 6, 8},
		12,
		"(1 1 1 1 1 1 1 1 1 1 1 1)(1 1 1 1 1 1 6)(1 1 1 1 8)(6 6)",
	)

	// act, assert
	testCombinationSumSolution(t, *entry)
}

func testCombinationSumSolution(t *testing.T, entry CombinationSumEntry) {
	// arrange
	exercise := backtracking.CombinationSum{}

	// act
	sln := exercise.Solution(entry.Input.A, entry.Input.Sum)

	// asserct
	if sln != entry.Output {
		t.Errorf("Expected %v but found %v in %v \n", entry.Output, sln, entry.Name)
		t.Fail()
	}
}
