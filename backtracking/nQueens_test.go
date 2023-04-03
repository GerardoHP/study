package backtracking_test

import (
	"study/backtracking"
	"study/utils"
	"testing"
)

type NQueensEntry struct {
	Test   string
	Input  int
	Output [][]int
}

func NewNQueensEntry(test string, input int, output [][]int) *NQueensEntry {
	return &NQueensEntry{
		Test:   test,
		Input:  input,
		Output: output,
	}
}

func TestNQueens_Solution_Test1(t *testing.T) {
	// arrange
	entry := NewNQueensEntry(
		"test 1",
		4,
		[][]int{
			{2, 4, 1, 3},
			{3, 1, 4, 2},
		})

	// act, assert
	testNQueens_Solution(t, entry)
}

func TestNQueens_Solution_Test5(t *testing.T) {
	// arrange
	entry := NewNQueensEntry(
		"test 5",
		6,
		[][]int{{2, 4, 6, 1, 3, 5},
			{3, 6, 2, 5, 1, 4},
			{4, 1, 5, 2, 6, 3},
			{5, 3, 1, 6, 4, 2}})

	// act, assert
	testNQueens_Solution(t, entry)
}

func testNQueens_Solution(t *testing.T, entry *NQueensEntry) {
	// arrange
	exercise := backtracking.NQueens{}

	// act
	sln := exercise.Solution(entry.Input)

	// assert
	if len(sln) != len(entry.Output) {
		t.Fatal()
	}

	for i := 0; i < len(entry.Output); i++ {
		a := sln[i]
		b := entry.Output[i]
		if !utils.SlicesEqual(a, b) {
			t.Errorf("expected %v, but found %v in %v \n", b, a, entry.Test)
			t.Fail()
		}
	}

}
