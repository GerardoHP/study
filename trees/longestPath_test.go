package trees_test

import (
	"study/trees"
	"testing"
)

type LongestPathEntry struct {
	Test   string
	Input  string
	Output int
}

func NewLongestPathEntry(test, input string, output int) *LongestPathEntry {
	return &LongestPathEntry{
		Test:   test,
		Input:  input,
		Output: output,
	}
}

var (
	longestPathEntryTest1 = NewLongestPathEntry(
		"Test 1",
		"user\f\tpictures\f\tdocuments\f\t\tnotes.txt",
		24,
	)
	longestPathEntryTest2 = NewLongestPathEntry(
		"Test 2",
		"user\f\tpictures\f\t\tphoto.png\f\t\tcamera\f\tdocuments\f\t\tlectures\f\t\t\tnotes.txt",
		33,
	)
	longestPathEntryTest3 = NewLongestPathEntry(
		"Test 3",
		"a",
		0,
	)
)

func TestLongestPath_Solution_Test_Table(t *testing.T) {
	testTable := []*LongestPathEntry{
		longestPathEntryTest1,
		longestPathEntryTest2,
		longestPathEntryTest3,
	}

	for _, test := range testTable {
		testLongestPathSolution(t, *test)
	}
}

func TestLongestPath_Solution_Test1(t *testing.T) {
	test := *longestPathEntryTest1
	testLongestPathSolution(t, test)
}

func TestLongestPath_Solution_Test2(t *testing.T) {
	test := *longestPathEntryTest2
	testLongestPathSolution(t, test)
}

func TestLongestPath_Solution_Test3(t *testing.T) {
	test := *longestPathEntryTest3
	testLongestPathSolution(t, test)
}

func testLongestPathSolution(t *testing.T, test LongestPathEntry) {
	// arrange
	exercise := trees.LongestPath{}

	// act
	sln := exercise.Solution(test.Input)

	// assert
	if sln != test.Output {
		t.Errorf("expected %v but found %v in %v \n", test.Output, sln, test.Test)
		t.Fail()
	}

}
