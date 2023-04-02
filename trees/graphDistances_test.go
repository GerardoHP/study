package trees_test

import (
	"study/trees"
	"study/utils"
	"testing"
)

type GraphDistancesEntry struct {
	Test   string
	Input  GraphDistancesInput
	Output []int
}

type GraphDistancesInput struct {
	g [][]int
	s int
}

func NewGraphDistanceEntry(test string, s int, output []int, g [][]int) *GraphDistancesEntry {
	return &GraphDistancesEntry{
		Test:   test,
		Output: output,
		Input: GraphDistancesInput{
			g: g,
			s: s,
		},
	}
}

// g = [[-1, 3, 2],
// [2, -1, 0],
// [-1, 0, -1]]
// and s = 0, the output should be
// solution(g, s) = [0, 2, 2]
var (
	graphDistanceEntryExtra = NewGraphDistanceEntry(
		"Test extra",
		0,
		[]int{0},
		[][]int{
			//A  B	 C	D 	E 	F 	G 	H
			{0, 10, 3, 8, -1, -1, -1, -1},
			{10, 0, 8, -1, 6, -1, -1, -1},
			{3, 8, 0, 4, -1, 9, -1, -1},
			{8, -1, 4, 0, -1, -1, 7, -1},
			{-1, 6, -1, -1, 0, 3, -1, 11},
			{-1, -1, 9, -1, 3, 0, 1, 8},
			{-1, -1, -1, 7, -1, 1, 0, 5},
			{-1, -1, -1, -1, 11, 8, 6, 0}})
	graphDistanceEntry1 = NewGraphDistanceEntry("Test 1", 0, []int{0, 2, 2}, [][]int{{-1, 3, 2}, {2, -1, 0}, {-1, 0, -1}})
)

func TestGraphDistances_Solution_Test_Table(t *testing.T) {
	testTable := []GraphDistancesEntry{
		*graphDistanceEntry1,
	}

	for _, test := range testTable {
		testGraphDistancesSolution(t, test)
	}
}

func TestGraphDistances_Solution_1(t *testing.T) {
	testGraphDistancesSolution(t, *graphDistanceEntry1)
}

func testGraphDistancesSolution(t *testing.T, entry GraphDistancesEntry) {
	// arrange
	exercise := trees.GraphDistances{}

	// act
	sln := exercise.Solution(entry.Input.g, entry.Input.s)

	// assert
	if !utils.SlicesEqual(sln, entry.Output) {
		t.Errorf("expected %v but found %v in %v \n", entry.Output, sln, entry.Test)
		t.Fail()
	}
}
