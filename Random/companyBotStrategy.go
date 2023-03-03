package Random

import (
	"fmt"
	"study/Interfaces"
)

type CompanyBotStrategy struct {
}

var _ Interfaces.Exercise = (*CompanyBotStrategy)(nil)

func (CompanyBotStrategy) solution(trainingData [][]int) float64 {
	sum := 0.0
	l := 0.0
	for _, t := range trainingData {
		fmt.Println(t[1])
		if t[1] == 1 {
			sum += float64(t[0])
			l++
		}
	}

	if sum == 0 {
		return 0
	}

	return float64(sum / l)
}

// Execute implements Interfaces.Exercise
func (c CompanyBotStrategy) Execute() {
	example := make([][]int, 0)
	example = append(example, []int{4, -1})
	example = append(example, []int{0, 0})
	example = append(example, []int{5, -1})

	x := c.solution(example)
	fmt.Println(x)
}
