package heaps_test

import (
	"testing"

	"study/heaps"
)

// nums: [7, 6, 5, 4, 3, 2, 1]
// k: 2
// expected: 6
type KthLargestElementInput struct {
	nums []int
	k    int
}

type entry struct {
	Input  KthLargestElementInput `json:"input"`
	Output int
	Test   string
}

func TestKthLargestElement_Solution(t *testing.T) {
	var tableTest []entry
	entry1 := NewKthLargestInput([]int{7, 6, 5, 4, 3, 2, 1}, 2, 6, "Test 1")
	entry7 := NewKthLargestInput([]int{3, 1, 2, 4}, 2, 3, "Test 7")
	tableTest = append(tableTest, *entry1)
	tableTest = append(tableTest, *entry7)
	exercise := heaps.KthLargestElement{}
	for _, test := range tableTest {
		sln := exercise.Solution(test.Input.nums, test.Input.k)
		if sln != test.Output {
			t.Errorf("expected %v, got %v in test %v", test.Output, sln, test.Test)
			t.Fail()
		}
	}
}

func NewKthLargestInput(nums []int, k, expected int, test string) *entry {
	return &entry{
		Input:  KthLargestElementInput{nums: nums, k: k},
		Output: expected,
		Test:   test,
	}
}
