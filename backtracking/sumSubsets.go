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
	sums := [][]int{}
	solveSumSubsets(arr, []int{}, num, &sums)
	return sums
}

func solveSumSubsets(arr, sum []int, num int, sums *[][]int) {
	for k, v := range arr {
		newSum := make([]int, len(sum))
		copy(newSum, sum)
		newSum = append(newSum, v)

		newArr := make([]int, len(arr)-1)
		copy(newArr[:k], arr[:k])
		copy(newArr[k:], arr[k+1:])

		if sumArray(newSum) == num && !exist(newSum, *sums) {
			*sums = append(*sums, newSum)
			break
		}

		solveSumSubsets(newArr, newSum, num, sums)
	}
}

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func exist(arr []int, arrs [][]int) bool {
	current := arrHash(arr)
	for _, v := range arrs {
		if len(v) == len(arr) && current == arrHash(v) {
			return true
		}
	}

	return false
}

func arrHash(arr []int) string {
	str := strings.Builder{}
	sort.Ints(arr)
	for _, v := range arr {
		str.WriteString(strconv.Itoa(v))
	}

	return str.String()
}
