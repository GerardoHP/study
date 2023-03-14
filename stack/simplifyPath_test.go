package stack_test

import (
	"study/stack"
	"testing"
)

type SimplifyPathInput struct {
	path string
}

type entry struct {
	Test   string
	Input  SimplifyPathInput `json:"input"`
	Output string            `json:"output"`
}

func TestSimplifyPath_Solution(t *testing.T) {
	// arrange
	tableTest := []*entry{
		NewSimplifyPathEntry("Test 1", "/home/a/./x/../b//c/", "/home/a/b/c"),
		NewSimplifyPathEntry("Test 3", "/../", "/"),
	}

	exercise := stack.SimplifyPath{}

	// act
	for _, test := range tableTest {
		sln := exercise.Solution(test.Input.path)

		// assert
		if sln != test.Output {
			t.Errorf("expected %v, got %v in test %v", test.Output, sln, test.Test)
			t.Fail()
		}
	}
}

func NewSimplifyPathEntry(test, path, output string) *entry {
	return &entry{
		Input:  SimplifyPathInput{path: path},
		Output: output,
		Test:   test,
	}
}
