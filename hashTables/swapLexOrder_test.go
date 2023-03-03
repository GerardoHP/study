package hashTables_test

import (
	"study/hashTables"
	"testing"
)

type SwapLexOrderInput struct {
	str   string  `json:"str"`
	pairs [][]int `json:"pairs"`
}

type entry struct {
	Input  SwapLexOrderInput `json:"input"`
	Output string            `json:"output"`
}

func Test1(t *testing.T) {
	// Arrange
	entry := newEntryFromData("abdc", [][]int{{1, 4}, {3, 4}}, "dbca")
	exercise := hashTables.SwapLexOrder{}

	// Act
	result := exercise.Solution(entry.Input.str, entry.Input.pairs)

	// Assert
	if result != entry.Output {
		t.Fail()
	}
}

// Input:
// str: ""
// pairs: [[1,4], [7,8]]
// Expected Output: "dbcaefhg"
func Test2(t *testing.T) {
	// Arrange
	entry := newEntryFromData("abcdefgh", [][]int{{1, 4}, {7, 8}}, "dbcaefhg")
	exercise := hashTables.SwapLexOrder{}

	// Act
	result := exercise.Solution(entry.Input.str, entry.Input.pairs)

	// Assert
	if result != entry.Output {
		t.Fail()
	}
}

// str: "acxrabdz"
// pairs:{{1,3}, {6,8},{3,8}, {2,7}}
// Expected Output:"zdxrabca"
func Test3(t *testing.T) {
	// Arrange
	entry := newEntryFromData("acxrabdz", [][]int{{1, 3}, {6, 8}, {3, 8}, {2, 7}}, "zdxrabca")
	exercise := hashTables.SwapLexOrder{}

	// Act
	result := exercise.Solution(entry.Input.str, entry.Input.pairs)

	// Assert
	if result != entry.Output {
		t.Fail()
	}
}

// str: "lvvyfrbhgiyexoirhunnuejzhesylojwbyatfkrv"
// pairs: {{13,23}, {13,28}, {15,20}, {24,29},{6,7}, {3,4}, {21,30}, {2,13}, {12,15},{19,23}, {10,19}, {13,14}, {6,16}, {17,25},{6,21}, {17,26}, {5,6}, {12,24}}
// Expected Output:"lyyvurrhgxyzvonohunlfejihesiebjwbyatfkrv"
func Test7(t *testing.T) {
	// Arrange
	entry := newEntryFromData("lvvyfrbhgiyexoirhunnuejzhesylojwbyatfkrv",
		[][]int{{13, 23}, {13, 28}, {15, 20}, {24, 29}, {6, 7}, {3, 4}, {21, 30}, {2, 13}, {12, 15}, {19, 23}, {10, 19}, {13, 14}, {6, 16}, {17, 25}, {6, 21}, {17, 26}, {5, 6}, {12, 24}},
		"lyyvurrhgxyzvonohunlfejihesiebjwbyatfkrv")
	exercise := hashTables.SwapLexOrder{}

	// Act
	result := exercise.Solution(entry.Input.str, entry.Input.pairs)

	// Assert
	if result != entry.Output {
		t.Fail()
	}
}

func newSwapLexOrderFromData(str string, pairs [][]int) *SwapLexOrderInput {
	return &SwapLexOrderInput{str: str, pairs: pairs}
}

func newEntryFromData(str string, pairs [][]int, expected string) *entry {
	return &entry{Input: *newSwapLexOrderFromData(str, pairs), Output: expected}
}
