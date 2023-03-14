package stack_test

import (
	"study/stack"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	// arrange
	s := stack.Stack{}
	s.Push(100)
	s.Push(200)
	s.Push(300)

	// act
	pop := s.Pop()

	// assert
	if pop != 300 {
		t.Fail()
	}
}
