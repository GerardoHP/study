package dynamicProgramming_test

import (
	"study/dynamicProgramming"
	"testing"
)

type FillingBlocksEntry struct {
	Name   string
	Input  int `json:"input"`
	Output int `json:"output"`
}

func NewFillingBlocksEntry(name string, input, output int) *FillingBlocksEntry {
	return &FillingBlocksEntry{
		Name:   name,
		Input:  input,
		Output: output,
	}
}

func TestFillingBlocks_Solution(t *testing.T) {
	// arrange
	testTable := make(map[int]*FillingBlocksEntry) // []*FillingBlocksEntry{}
	testTable[2] = NewFillingBlocksEntry("test 2", 4, 36)

	// act, assert
	for k, v := range testTable {
		t.Logf("testing %v \n", k)
		testFillingBlocksSolution(t, v)
	}
}

func testFillingBlocksSolution(t *testing.T, entry *FillingBlocksEntry) {
	// arrange
	exercise := dynamicProgramming.FillingBlocks{}

	// act
	sln := exercise.Solution(entry.Input)

	// assert
	if sln != entry.Output {
		t.Errorf("expected %v, but found %v in %v \n", entry.Output, sln, entry.Name)
		t.Fail()
	}
}
