package backtracking

import (
	"study/interfaces"
)

type ClimbingStaircase struct{}

func (b ClimbingStaircase) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*ClimbingStaircase)(nil)

func (ClimbingStaircase) Solution(n int, k int) [][]int {
	paths := [][]int{}
	backtracking(k, n, []int{}, &paths)
	return paths
}

func backtracking(k, remainingSteps int, path []int, paths *[][]int) {
	if remainingSteps == 0 {
		*paths = append(*paths, path)
		return
	}

	for i := 1; i <= k; i++ {
		if remainingSteps >= i {
			newPath := make([]int, len(path))
			copy(newPath, path)
			newPath = append(newPath, i)
			backtracking(k, remainingSteps-i, newPath, paths)
		}
	}
}
