package trees_test

import (
	"encoding/json"
	"study/trees"
	"study/utils"
	"testing"
)

type LargestValuesInTreeRowsEntry struct {
	Test   string
	Input  treeEntry `json:"input"`
	Output []int     `json:"output"`
}

type treeEntry struct {
	T *trees.Tree `json:"t"`
}

func NewLargestValuesInTreeRowsEntry(test, jsonStr string) *LargestValuesInTreeRowsEntry {
	var entry LargestValuesInTreeRowsEntry
	json.Unmarshal([]byte(jsonStr), &entry)
	entry.Test = test
	return &entry
}

var (
	test1 = NewLargestValuesInTreeRowsEntry("Test 1", "{\n    \"input\": {\n        \"t\": {\n            \"value\": -1,\n            \"left\": {\n                \"value\": 5,\n                \"left\": null,\n                \"right\": null\n            },\n            \"right\": {\n                \"value\": 7,\n                \"left\": null,\n                \"right\": {\n                    \"value\": 1,\n                    \"left\": null,\n                    \"right\": null\n                }\n            }\n        }\n    },\n    \"output\": [\n        -1,\n        7,\n        1\n    ]\n}")
	// test1 = NewLargestValuesInTreeRowsEntry("test", "{\n  \"test\": 1,\n  \"input\": {\n    \"t\": {\n      \"value\": 5,\n      \"left\": {\n        \"value\": 2,\n        \"left\": {\n          \"value\": 1,\n          \"left\": null,\n          \"right\": null\n        },\n        \"right\": {\n          \"value\": 3,\n          \"left\": null,\n          \"right\": null\n        }\n      },\n      \"right\": {\n        \"value\": 6,\n        \"left\": null,\n        \"right\": {\n          \"value\": 8,\n          \"left\": {\n            \"value\": 7,\n            \"left\": null,\n            \"right\": null\n          },\n          \"right\": null\n        }\n      }\n    },\n    \"queries\":[4, 5, 6]\n  },\n  \"output\": {\n    \"value\": 3,\n    \"left\": {\n      \"value\": 2,\n      \"left\": {\n        \"value\": 1,\n        \"left\": null,\n        \"right\": null\n      },\n      \"right\": null\n    },\n    \"right\": {\n      \"value\": 8,\n      \"left\": {\n        \"value\": 7,\n        \"left\": null,\n        \"right\": null\n      },\n      \"right\": null\n    }\n  }\n}")
)

func TestTableLargestValuesInTreeRows_Solution(t *testing.T) {
	testTable := []LargestValuesInTreeRowsEntry{
		*test1,
	}

	for _, test := range testTable {
		testLargestValuesInTreeRowsSolution(t, test)
	}
}

func TestLargestValuesInTreeRows_Solution(t *testing.T) {
	testLargestValuesInTreeRowsSolution(t, *test1)
}

func testLargestValuesInTreeRowsSolution(t *testing.T, test LargestValuesInTreeRowsEntry) {
	// arrange
	x := trees.LargestValuesInTreeRows{}

	// act
	sln := x.Solution(test.Input.T)

	// assert
	if !utils.SlicesEqual(sln, test.Output) {
		t.Errorf("expected %v, but found %v in %v \n", test.Output, sln, test.Test)
		t.Fail()
	}
}
