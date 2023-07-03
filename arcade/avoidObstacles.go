package arcade

import "fmt"

type AvoidObstacles struct{}

func (AvoidObstacles) Solution(inputArray []int) int {
	c := 0
	max := 0
	m := make(map[int]bool)
	for _, v := range inputArray {
		m[v] = true
		if v > max {
			max = v
		}
	}

	fmt.Println(m)
	fmt.Println(max)
	for i := 0; i <= max; i++ {
		_, found := m[i]
		fmt.Println(c)
		if !found {
			c++
		}

		if found && i < max {
			c = 0
		}
	}

	return c
}
