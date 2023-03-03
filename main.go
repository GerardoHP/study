package main

import (
	"study/Interfaces"
	"study/Lists"
	"study/Random"
	"study/Trees"
)

type ExerciseType int64

const (
	simpleLinkedList ExerciseType = iota
	isListPalindrome
	hasPathWithGivenSum
	isTreeSymmetric
	findProfession
	treeToList
	isSubTree
	restoreBinaryTree
	companyBotStrategy
	findSubstrings
)

var current ExerciseType = ExerciseType(9)

func main() {
	var exercise Interfaces.Exercise
	switch current {
	case isListPalindrome:
		exercise = Lists.IsListPalindrome{}
	case hasPathWithGivenSum:
		exercise = Trees.HasPathWithGivenSum{}
	case isTreeSymmetric:
		exercise = Trees.IsTreeSymmetric{}
	case findProfession:
		exercise = Trees.FindProfession{}
	case treeToList:
		exercise = Trees.TreeToList{}
	case isSubTree:
		exercise = Trees.IsSubTree{}
	case restoreBinaryTree:
		exercise = Trees.RestoreBinaryTree{}
	case companyBotStrategy:
		exercise = Random.CompanyBotStrategy{}
	case findSubstrings:
		exercise = Trees.FindSubstrings{}
	default:
		exercise = nil
	}

	exercise.Execute()
}
