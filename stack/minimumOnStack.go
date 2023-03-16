package stack

import (
	"math"
	"strconv"
	"strings"
	"study/Interfaces"
)

type MinimumOnStack struct{}

func (m MinimumOnStack) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ Interfaces.Exercise = (*MinimumOnStack)(nil)

func (MinimumOnStack) Solution(operations []string) []int {
	return solution(operations)
}

func (s Stack) Min() int {
	min := math.MaxInt
	for _, v := range s {
		if v.(int) < min {
			min = v.(int)
		}
	}

	return min
}

func solution(operations []string) []int {
	mins := []int{}
	s := Stack{}
	for _, operation := range operations {
		switch {
		case strings.HasPrefix(operation, "push"):
			// push
			pushStr := strings.ReplaceAll(operation, "push ", "")
			push, _ := strconv.Atoi(pushStr)
			s.Push(push)
		case strings.HasPrefix(operation, "pop"):
			// pop
			s.Pop()
		case strings.HasPrefix(operation, "min"):
			// min
			mins = append(mins, s.Min())
		}
	}

	return mins
}
