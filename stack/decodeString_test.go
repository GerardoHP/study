package stack_test

import (
	"study/stack"
	"testing"
)

type decodeStringEntry struct {
	Test   string
	Input  string
	Output string
}

func TestDecodeString_Solution_Test1(t *testing.T) {
	test := NewDecodeStringEntry("Test 1", "4[ab]", "abababab")
	testDecodeStringSolution(t, test)
}

func TestDecodeString_Solution_Test3(t *testing.T) {
	test := NewDecodeStringEntry("Test 3", "z1[y]zzz2[abc]", "zyzzzabcabc")
	testDecodeStringSolution(t, test)
}

func TestDecodeString_Solution_Test10(t *testing.T) {
	test := NewDecodeStringEntry(
		"Test 10",
		"100[codesignal]",
		"codesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignalcodesignal",
	)
	testDecodeStringSolution(t, test)
}

func testDecodeStringSolution(t *testing.T, test *decodeStringEntry) {
	// arrange
	exercise := stack.DecodeString{}

	// act
	result := exercise.Solution(test.Input)

	// assert
	if result != test.Output {
		t.Errorf("expected %v, got %v in test %v", test.Output, result, test.Test)
		t.Fail()
	}
}

func NewDecodeStringEntry(test, input, output string) *decodeStringEntry {
	return &decodeStringEntry{
		Test:   test,
		Input:  input,
		Output: output,
	}
}
