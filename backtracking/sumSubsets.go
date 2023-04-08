package backtracking

import (
	"sort"
	"strconv"
	"strings"
	"study/interfaces"
)

type SumSubsets struct{}

func (s SumSubsets) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*SumSubsets)(nil)

func (SumSubsets) Solution(arr []int, num int) [][]int {
	if len(arr) == 0 || num == 0 {
		return [][]int{[]int{}}
	}

	sums := [][]int{}
	hashes := make(map[string]bool)
	solveSumSubsets(arr, []int{}, num, &sums, hashes)

	return sums
}

func solveSumSubsets(arr, sum []int, num int, sums *[][]int, hashes map[string]bool) {
	for k, v := range arr {
		newSum := make([]int, len(sum))
		copy(newSum, sum)
		newSum = append(newSum, v)

		newArr := make([]int, len(arr)-1)
		copy(newArr[:k], arr[:k])
		copy(newArr[k:], arr[k+1:])

		s, currentHash := sumArray(newSum)
		if _, found := hashes[currentHash]; found || s >= num {
			if !found && s == num {
				hashes[currentHash] = true
				*sums = append(*sums, newSum)
			}

			break
		}

		solveSumSubsets(newArr, newSum, num, sums, hashes)
	}
}

func sumArray(arr []int) (int, string) {
	sort.Ints(arr)
	sum := 0
	str := strings.Builder{}
	for _, v := range arr {
		sum += v
		str.WriteString(strconv.Itoa(v))
	}

	return sum, str.String()
}
