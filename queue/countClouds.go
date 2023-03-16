package queue

import "study/Interfaces"

type CountClouds struct {
}

func (c CountClouds) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ Interfaces.Exercise = (*CountClouds)(nil)

func (CountClouds) Solution(skyMap [][]string) int {
	if len(skyMap) == 0 || len(skyMap[0]) == 0 {
		return 0
	}

	c := 0
	for i, v1 := range skyMap {
		for j, v2 := range v1 {
			if v2 == "1" {
				c++
			}

			navigate(i, j, skyMap)
		}
	}

	return c
}

func navigate(i, j int, m [][]string) {
	if m[i][j] != "1" {
		return
	}

	m[i][j] = "*"
	if i > 0 {
		navigate(i-1, j, m)
	}

	if i < len(m)-1 {
		navigate(i+1, j, m)
	}

	if j > 0 {
		navigate(i, j-1, m)
	}

	if j < len(m[i])-1 {
		navigate(i, j+1, m)
	}
}
