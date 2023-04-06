package backtracking

import (
	"study/interfaces"
)

type NQueens struct{}

func (N NQueens) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*NQueens)(nil)

func (NQueens) Solution(n int) [][]int {
	paths := [][]int{}
	solveNQueens(n, []int{}, &paths)
	return paths
}

func solveNQueens(n int, path []int, paths *[][]int) {
	if len(path) == n {
		if pathComplies(path) {
			*paths = append(*paths, path)
		}

		return
	}

	for i := 0; i < n; i++ {
		newPath := make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, i+1)
		solveNQueens(n, newPath, paths)
	}
}

func pathComplies(path []int) bool {
	if !rowComplies(path) || !diagonalComplies(path) {
		return false
	}

	return true
}

func diagonalComplies(path []int) bool {
	for i, v := range path {
		l := len(path)
		for j := i + 1; j < l; j++ {
			if v+j-i == path[j] || v-j+i == path[j] {
				return false
			}
		}
	}

	return true
}

func rowComplies(path []int) bool {
	m := make(map[int]bool)
	for _, v := range path {
		if _, found := m[v]; !found {
			m[v] = true
		} else {
			return false
		}
	}

	return true
}

func fixedNextPrev(l, j, current int) (next, prev int) {
	next = current + j
	prev = current - j
	if next >= l {
		next -= l
	}

	if prev < 0 {
		prev += l
	}

	return
}

func fixedValue(l, j, v int) (up, down int) {
	up = v + j
	down = v - j
	if up > l {
		up -= l
	}

	if down < 1 {
		down += l
	}

	return
}
