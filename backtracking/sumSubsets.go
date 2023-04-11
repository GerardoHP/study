package backtracking

import (
	"hash/maphash"
	"sort"
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
	hashes := make(map[uint64]bool)
	solveSumSubsets(arr, []int{}, num, 0, &sums, hashes, maphash.MakeSeed())

	return sums
}

func solveSumSubsets(nums, current []int, target int, index int, sums *[][]int, hashes map[uint64]bool, seed maphash.Seed) {
	if target == 0 {
		currentHash := getHash(current, seed)
		if _, found := hashes[currentHash]; !found {
			temp := make([]int, len(current))
			copy(temp, current)
			*sums = append(*sums, temp)
			hashes[currentHash] = true
		}

		return
	}

	if index == len(nums) {
		return
	}

	if nums[index] <= target {
		current = append(current, nums[index])
		solveSumSubsets(nums, current, target-nums[index], index+1, sums, hashes, seed)
		current = current[:len(current)-1]
	}

	solveSumSubsets(nums, current, target, index+1, sums, hashes, seed)
}

func getHash(arr []int, seed maphash.Seed) uint64 {
	var hash maphash.Hash
	hash.SetSeed(seed)
	sort.Ints(arr)
	for _, v := range arr {
		hash.WriteByte(byte(v))
	}

	return hash.Sum64()
}
