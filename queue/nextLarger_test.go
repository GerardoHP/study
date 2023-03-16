package queue_test

import (
	"study/queue"
	"study/utils"
	"testing"
)

type nextLargerEntry struct {
	Test   string
	Input  []int
	Output []int
}

func NewNextLargerEntry(test string, input, output []int) *nextLargerEntry {
	return &nextLargerEntry{
		Test:   test,
		Input:  input,
		Output: output,
	}
}

var (
	test1 = NewNextLargerEntry("Test 1", []int{6, 7, 3, 8}, []int{7, 8, 8, -1})
	tests = map[int]*nextLargerEntry{
		1: test1,
	}
)

func TestNextLarger_Solution_1(t *testing.T) {
	test := tests[1]
	testNextLargerSolution(t, *test)
}

func Teste_All(t *testing.T) {
	for _, test := range tests {
		testNextLargerSolution(t, *test)
	}
}

func testNextLargerSolution(t *testing.T, entry nextLargerEntry) {
	// arrange
	exercise := queue.NextLarger{}

	// act
	sln := exercise.Solution(entry.Input)

	// assert
	if !utils.SlicesEqual(sln, entry.Output) {
		t.Errorf("expected %v, found %v in %v \n", entry.Output, sln, entry.Test)
		t.Fail()
	}
}
