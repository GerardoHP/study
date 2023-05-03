package dynamicProgramming

import (
	"fmt"
	"study/interfaces"
)

type FillingBlocks struct{}

var _ interfaces.Exercise = (*FillingBlocks)(nil)

func (f FillingBlocks) Execute() {
	//TODO implement me
	panic("implement me")
}

func (FillingBlocks) Solution(n int) int {
	if n == 1 {
		return 1
	}

	matrix := make([][][][]int, n+1)
	initN4Matrix(n, &matrix)
	matrix[0][0][0][0] = 1
	for i := 1; i <= n; i++ {
		for a := range matrix[i] {
			for b := range matrix[i][a] {
				for c := range matrix[i][a][b] {
					matrix[i][a][b][c] = matrix[i-1][a][b][1-c]
					if i >= 2 {
						matrix[i][a][b][c] += matrix[i-2][a][1-b][c] + matrix[i-2][1-a][b][c]
					}
				}
			}
		}
	}

	return matrix[n][0][0][0]
}

func initN4Matrix(n int, matrix *[][][][]int) {
	for i := range *matrix {
		(*matrix)[i] = make([][][]int, 2)
		for a := range (*matrix)[i] {
			(*matrix)[i][a] = make([][]int, 2)
			for b := range (*matrix)[i][a] {
				(*matrix)[i][a][b] = make([]int, 2)
			}
		}
	}

	fmt.Println(matrix)
}
