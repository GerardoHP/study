package arcade_test

import (
	"study/arcade"
	"testing"
)

type AvoidObstaclesEntry struct {
	Input  []int `json:"input"`
	Output int   `json:"output"`
	Name   string
}

func NewAvoidObstacleEntry(input []int, output int, name string) *AvoidObstaclesEntry {
	return &AvoidObstaclesEntry{
		Name:   name,
		Input:  input,
		Output: output,
	}
}

func TestAvoidObstacles_Solution_Table_Tests(t *testing.T) {
	// arrange
	table := []*AvoidObstaclesEntry{}
	entry1 := NewAvoidObstacleEntry([]int{5, 3, 6, 7, 9}, 4, "test 1")
	table = append(table, entry1)

	// act, assert
	for _, entry := range table {
		testAvoidObstacles(t, entry)
	}
}

func testAvoidObstacles(t *testing.T, entry *AvoidObstaclesEntry) {
	// arrange
	exercise := arcade.AvoidObstacles{}

	//act
	sln := exercise.Solution(entry.Input)

	// assert
	if sln != entry.Output {
		t.Errorf("expected %v, but found %v in %v \n", entry.Output, sln, entry.Name)
		t.Fail()
	}
}
