package main

import (
	"study/Random"
	"study/Trees"
	"study/heaps"
	interfaces "study/interfaces"
	"study/lists"
	"study/queue"
	"study/stack"
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
	kthLargestElement
	countClouds
	nextLarger
	decodeString
	minimumOnStack
	simplifyPath
)

var current ExerciseType = ExerciseType(9)

func main() {
	var exercise interfaces.Exercise
	switch current {
	case isListPalindrome:
		exercise = lists.IsListPalindrome{}
	case hasPathWithGivenSum:
		exercise = trees.HasPathWithGivenSum{}
	case isTreeSymmetric:
		exercise = trees.IsTreeSymmetric{}
	case findProfession:
		exercise = trees.FindProfession{}
	case treeToList:
		exercise = trees.TreeToList{}
	case isSubTree:
		exercise = trees.IsSubTree{}
	case restoreBinaryTree:
		exercise = trees.RestoreBinaryTree{}
	case companyBotStrategy:
		exercise = random.CompanyBotStrategy{}
	case findSubstrings:
		exercise = trees.FindSubstrings{}
	case kthLargestElement:
		exercise = heaps.KthLargestElement{}
	case countClouds:
		exercise = queue.CountClouds{}
	case nextLarger:
		exercise = queue.NextLarger{}
	case decodeString:
		exercise = stack.DecodeString{}
	case minimumOnStack:
		exercise = stack.MinimumOnStack{}
	case simplifyPath:
		exercise = stack.SimplifyPath{}
	default:
		exercise = nil
	}

	exercise.Execute()
}
