package dynamicProgramming

import (
	"fmt"
)

var m = make(map[int]int)

func Execute() {
	x := solution(5)
	y := solution(10)
	fmt.Println(x)
	fmt.Println(y)
}

func solution(n int) int {
	result := 0

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if r, found := m[n]; found {
		return r
	}

	result = solution(n-1) + solution(n-2)
	m[n] = result
	return result
}
