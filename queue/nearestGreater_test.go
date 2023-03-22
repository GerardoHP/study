package queue_test

import (
	"study/queue"
	"study/utils"
	"testing"
)

type NearestGreaterEntry struct {
	Test   string
	Input  []int
	Output []int
}

func NewNearestGreaterEntry(test string, a, b []int) *NearestGreaterEntry {
	return &NearestGreaterEntry{
		Test:   test,
		Input:  a,
		Output: b,
	}
}

func TestNearestGreater_Solution_Test1(t *testing.T) {
	test := NewNearestGreaterEntry("Test 1", []int{1, 4, 2, 1, 7, 6}, []int{1, 4, 1, 2, -1, 4})
	testNearestGreaterSolution(t, *test)
}

func testNearestGreaterSolution(t *testing.T, test NearestGreaterEntry) {
	// arrange
	exercise := queue.NearestGreater{}

	// act
	sln := exercise.Solution(test.Input)

	// assert
	if !utils.SlicesEqual(sln, test.Output) {
		t.Errorf("expected %v, but found %v in %v \n ", test.Output, sln, test.Test)
		t.Fail()
	}
}
