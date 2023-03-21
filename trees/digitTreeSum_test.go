package trees_test

import (
	"encoding/json"
	"study/trees"
	"testing"
)

type DigitTreeSumEntry struct {
	Test   string
	Input  *trees.Tree
	Output int
}

func NewDigitTreeSumEntry(test string, output int, treeJson string) *DigitTreeSumEntry {
	var t trees.Tree
	json.Unmarshal([]byte(treeJson), &t)
	return &DigitTreeSumEntry{
		Test:   test,
		Input:  &t,
		Output: output,
	}
}

var (
	treeSumEntry1 = NewDigitTreeSumEntry("Test 1", 218, "{\n    \"value\": 1,\n    \"left\": {\n        \"value\": 0,\n        \"left\": {\n            \"value\": 3,\n            \"left\": null,\n            \"right\": null\n        },\n        \"right\": {\n            \"value\": 1,\n            \"left\": null,\n            \"right\": null\n        }\n    },\n    \"right\": {\n        \"value\": 4,\n        \"left\": null,\n        \"right\": null\n    }\n}")
)

func TestDigitTreeSum_Solution_Table(t *testing.T) {
	testTable := []DigitTreeSumEntry{
		*treeSumEntry1,
	}

	for _, test := range testTable {
		testDigitTreeSumSolution(t, test)
	}
}

func TestDigitTreeSum_Solution_Test1(t *testing.T) {
	testDigitTreeSumSolution(t, *treeSumEntry1)
}

func testDigitTreeSumSolution(t *testing.T, test DigitTreeSumEntry) {
	// arrange
	exercise := trees.DigitTreeSum{}

	// act
	sln := exercise.Solution(test.Input)

	// assert
	if sln != test.Output {
		t.Errorf("expected %v but found %v in %v \n", test.Output, sln, test.Test)
		t.Fail()
	}
}
