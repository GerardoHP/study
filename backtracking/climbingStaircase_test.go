package backtracking_test

import (
	"encoding/json"
	"io"
	"os"
	"study/backtracking"
	"study/utils"
	"testing"
)

type ClimbingStaircaseEntry struct {
	Input  ClimbingStaircaseInput `json:"input"`
	Output [][]int                `json:"output"`
}

type ClimbingStaircaseInput struct {
	N int `json:"n"`
	K int `json:"k"`
}

func NewClimbingStaircaseEntry(file string) (*ClimbingStaircaseEntry, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var entry ClimbingStaircaseEntry
	json.Unmarshal(byteValue, &entry)
	return &entry, nil
}

func DeconstructInput(input ClimbingStaircaseInput) (int, int) {
	return input.N, input.K
}

func TestClimbingStaircase_Solution(t *testing.T) {
	// arrange
	expectedCount := 5
	exercise := backtracking.ClimbingStaircase{}

	// act
	sln := exercise.Solution(4, 2)

	// assert
	if l := len(sln); l != expectedCount {
		t.Errorf("expected %v, but found %v \n", expectedCount, l)
		t.Fail()
	}
}

func TestClimbingStaircase_Solution2(t *testing.T) {
	// arrange
	entry, err := NewClimbingStaircaseEntry("testFiles/climbingStaircase_test-4.json")
	if err != nil {
		t.Fatal(err)
	}

	exercise := backtracking.ClimbingStaircase{}
	n, k := DeconstructInput(entry.Input)

	// act
	sln := exercise.Solution(n, k)

	// assert
	for k := range sln {
		a := sln[k]
		b := entry.Output[k]
		if !utils.SlicesEqual(a, b) {
			t.Errorf("expected %v, but found %v in array %v ", b, a, k)
			t.Fail()
		}
	}
}
