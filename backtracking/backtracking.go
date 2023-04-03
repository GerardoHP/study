package backtracking

import "study/interfaces"

type Backtracking struct{}

func (b Backtracking) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*Backtracking)(nil)

func (Backtracking) Solution(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, &result, []int{})
	return result
}

func backtrack(nums []int, result *[][]int, path []int) {
	// Si no quedan elementos por agregar al path, lo agregamos al resultado
	if len(nums) == 0 {
		*result = append(*result, path)
		return
	}

	// Iteramos sobre los elementos que quedan por agregar al path
	for i, num := range nums {
		// Creamos una copia del path y agregamos el elemento actual
		newPath := append(path, num)

		// Creamos una copia del slice de números y lo modificamos para no volver a usar el elemento actual
		newNums := make([]int, len(nums)-1)
		copy(newNums[:i], nums[:i])
		copy(newNums[i:], nums[i+1:])

		// Llamamos recursivamente a backtrack con el nuevo path y los números restantes
		backtrack(newNums, result, newPath)
	}
}
