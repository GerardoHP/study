package queue_test

import (
	"study/queue"
	"testing"
)

type CountCloudsEntry struct {
	Test   string
	Input  [][]string
	Output int
}

func NewCountCloudsEntry(test string, input [][]string, output int) *CountCloudsEntry {
	return &CountCloudsEntry{
		Test:   test,
		Input:  input,
		Output: output,
	}
}

func TestCountClouds_Solution(t *testing.T) {
	test := NewCountCloudsEntry("Test 1", [][]string{{"0", "1", "1", "0", "1"}, {"0", "1", "1", "1", "1"}, {"0", "0", "0", "0", "1"}, {"1", "0", "0", "1", "1"}}, 2)
	testCountCloudsSolution(t, *test)
}

func testCountCloudsSolution(t *testing.T, test CountCloudsEntry) {
	// arrange
	exercise := queue.CountClouds{}

	// act
	sln := exercise.Solution(test.Input)

	// assert
	if sln != test.Output {
		t.Errorf("expected %v, but found %v in %v \n /n", test.Output, sln, test.Test)
		t.Fail()
	}
}
