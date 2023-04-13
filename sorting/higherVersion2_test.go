package sorting_test

import (
	"study/sorting"
	"testing"
)

type HigherVersionInput struct {
	Ver1 string `json:"ver1"`
	Ver2 string `json:"ver2"`
}

type HigherVersionEntry struct {
	Name   string
	Input  HigherVersionInput `json:"input"`
	Output int                `json:"output"`
}

func NewHigherVersionEntry(name, ver1, ver2 string, output int) *HigherVersionEntry {
	return &HigherVersionEntry{
		Name: name,
		Input: HigherVersionInput{
			Ver1: ver1,
			Ver2: ver2,
		},
		Output: output,
	}
}

func testHigherVersion2Solution(t *testing.T, entry HigherVersionEntry) {
	// arrange
	exercise := sorting.HigherVersion2{}

	// act
	sln := exercise.Solution(entry.Input.Ver1, entry.Input.Ver2)

	// assert
	if entry.Output != sln {
		t.Errorf("expected %v, but found %v in %v \n", entry.Output, sln, entry.Name)
		t.Fail()
	}
}

func TestHigherVersion2_Solution_1(t *testing.T) {
	// arrange
	entry := NewHigherVersionEntry("Test 1", "1.2.2", "1.2.0", 1)

	// act, assert
	testHigherVersion2Solution(t, *entry)
}
