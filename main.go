package main

import (
	"study/Interfaces"
	"study/Lists"
)

type ExerciseType int64

const (
	SimpleLinkedList ExerciseType = iota
	isListPalindrome
)

var current ExerciseType = isListPalindrome

func main() {
	var exercise Interfaces.Exercise
	switch current {
	case isListPalindrome:
		exercise = Lists.IsListPalindrome{}
	default:
		exercise = nil
	}

	exercise.Execute()
}
