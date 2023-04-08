package backtracking_test

import (
	"encoding/json"
	"io"
	"os"
	"study/backtracking"
	"study/utils"
	"testing"
)

type SumSubsetsEntry struct {
	Input  SumSubsetsInput `json:"input"`
	Output [][]int         `json:"output"`
}

type SumSubsetsInput struct {
	Arr []int `json:"arr"`
	Num int   `json:"num"`
}

func NewSumsetsEntry(file string) (*SumSubsetsEntry, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var entry SumSubsetsEntry
	json.Unmarshal(byteValue, &entry)
	return &entry, nil
}

func TestNQueens_Solution_1(t *testing.T) {
	// arrange
	entry := SumSubsetsEntry{
		Input: SumSubsetsInput{
			Arr: []int{1, 2, 3, 4, 5},
			Num: 5,
		},
		Output: [][]int{
			{1, 4},
			{2, 3},
			{5},
		},
	}

	// act, assert
	testSumSubsets_Solution(t, entry)
}

func TestSumSubsets_Solution_6(t *testing.T) {
	// arrange
	entry, err := NewSumsetsEntry("testFiles/sumSubsets_test-6.json")

	// act, assert
	if err != nil {
		t.Fatal(err)
	}

	testSumSubsets_Solution(t, *entry)
}

func TestSumSubsets_Solution_10(t *testing.T) {
	// arrange
	entry, err := NewSumsetsEntry("testFiles/sumSubsets_test-10.json")

	// act, assert
	if err != nil {
		t.Fatal(err)
	}

	testSumSubsets_Solution(t, *entry)
}

func testSumSubsets_Solution(t *testing.T, entry SumSubsetsEntry) {
	// arrange
	arr := entry.Input.Arr
	num := entry.Input.Num
	r := entry.Output
	exercise := backtracking.SumSubsets{}

	// act
	sln := exercise.Solution(arr, num)

	// assert
	for k := range sln {
		a := sln[k]
		b := r[k]
		if len(a) != len(b) || !utils.SlicesEqual(a, b) {
			t.Errorf("expected %v, but found %v in test 1", b, a)
			t.Fail()
		}
	}
}
