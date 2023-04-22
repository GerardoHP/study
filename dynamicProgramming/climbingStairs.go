package dynamicProgramming

type ClimbingStairs struct{}

var m = make(map[int]int)

func (c ClimbingStairs) Solution(n int) int {
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

	result = c.Solution(n-1) + c.Solution(n-2)
	m[n] = result
	return result
}
