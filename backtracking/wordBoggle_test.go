package backtracking_test

import (
	"encoding/json"
	"io"
	"os"
	"study/backtracking"
	"study/utils"
	"testing"
)

type WordBoggleEntry struct {
	Name   string
	Input  WordBoggleInput `json:"input"`
	Output []string        `json:"output"`
}

type WordBoggleInput struct {
	Board [][]string `json:"board"`
	Words []string   `json:"words"`
}

func testWordBoggleSolution(t *testing.T, entry WordBoggleEntry) {
	// arrange
	exercise := backtracking.WordBoggle{}

	// act
	sln := exercise.Solution(entry.Input.Board, entry.Input.Words)

	// assert
	if len(sln) != len(entry.Output) {
		t.Fatal()
	}

	if !utils.StringSlicesEqual(sln, entry.Output) {
		t.Errorf("expected %v but found %v in %v \n", entry.Output, sln, entry.Name)
		t.Fail()
	}
}

func NewWordBoggleEntryFromFile(file string) (*WordBoggleEntry, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var entry WordBoggleEntry
	json.Unmarshal(bytes, &entry)
	return &entry, nil
}

func TestWordBoggle_Solution_1(t *testing.T) {
	entry := WordBoggleEntry{
		Name: "Test 1",
		Input: WordBoggleInput{
			Board: [][]string{
				{"R", "L", "D"},
				{"U", "O", "E"},
				{"C", "S", "O"},
			},
			Words: []string{
				"CODE",
				"SOLO",
				"RULES",
				"COOL"},
		},
		Output: []string{"CODE",
			"RULES"},
	}

	testWordBoggleSolution(t, entry)
}
