package dynamicProgramming

import (
	"fmt"
	"math"
)

func Execute() {
	//x := solution(5)
	y := solution(10)
	//fmt.Println(x)
	fmt.Println(y)
}

var (
	allPaths   = make(map[int][][]int)
	currentMax = math.MinInt
)

func solution(n int) int {
	path := []int{}
	paths := new([][]int)
	s := 0
	if p2, found := allPaths[currentMax]; found && n > currentMax {
		p := p2[0]
		for _, v := range p {
			path = append(path, v)
		}

		s = currentMax
	}

	solve(s, n, []int{1, 2}, path, paths)
	allPaths[n] = *paths
	if n > currentMax {
		currentMax = n
	}

	return len(*paths)
}

func solve(s, n int, comb, path []int, paths *[][]int) {
	if s >= n {
		if s == n {
			*paths = append(*paths, path)
		}

		return
	}

	for _, v := range comb {
		newPath := make([]int, len(path))
		newPath = append(path, v)

		if s+v <= n {
			solve(s+v, n, comb, newPath, paths)
		}
	}
}

func sum(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return a[0] + sum(a[1:])
}
