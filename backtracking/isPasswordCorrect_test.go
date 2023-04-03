package backtracking_test

import (
	"study/backtracking"
	"testing"
)

type IsPasswordCorrectEntry struct {
	Input  IsPasswordCorrectInput `json:"input"`
	Output bool                   `json:"output"`
}

type IsPasswordCorrectInput struct {
	Password string `json:"password"`
	Nums     []int  `json:"nums"`
}

func NewIsPasswordCorrectEntry(pass string, nums []int, output bool) *IsPasswordCorrectEntry {
	return &IsPasswordCorrectEntry{
		Input:  IsPasswordCorrectInput{Password: pass, Nums: nums},
		Output: output,
	}
}

func TestIsPasswordCorrect_Execute(t *testing.T) {
	// arrange
	exercise := backtracking.IsPasswordCorrect{}
	entry := NewIsPasswordCorrectEntry("132", []int{2, 3, 1}, true)

	// act
	sln := exercise.Solution(entry.Input.Password, entry.Input.Nums)

	// assert
	if sln != entry.Output {
		t.Errorf("expected %v, but found %v \n", entry.Output, sln)
		t.Fail()
	}
}
