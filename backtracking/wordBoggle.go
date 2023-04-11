package backtracking

import (
	"fmt"
	"study/interfaces"
)

type WordBoggle struct{}

type coordinate struct {
	x int
	y int
}

func (w WordBoggle) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*WordBoggle)(nil)

func (WordBoggle) Solution(board [][]string, words []string) []string {
	result := []string{}

	solveWordBoggle(board, make(map[string]bool), words[0], "", 0, 0)
	return result
}

func solveWordBoggle(board [][]string, visited map[string]bool, word, currentWord string, x, y int) bool {
	if word == currentWord {
		return true
	}

	if len(word) == len(currentWord) {
		return false
	}

	currentHash := getVertexHash(x, y)
	visited[currentHash] = true
	for i := x; i < len(board); i++ {
		for j := y; j < len(board[i]); j++ {
			coordinates := getNeighbors(i, j, len(board[i]), len(board), visited)
			fmt.Println(i, j, coordinates)
		}
	}

	return false
}

func getVertexHash(x, y int) string {
	return fmt.Sprintf("%v_%v", x, y)
}

func getNeighbors(x, y, width, height int, path map[string]bool) []coordinate {
	coordinates := []coordinate{}

	nextX := x + 1
	prevX := x - 1
	nextY := y + 1
	prevY := y - 1

	if prevX > -1 {
		coordinates = append(coordinates, coordinate{
			x: prevX,
			y: y,
		})
	}

	if nextX < height {
		coordinates = append(coordinates, coordinate{
			x: nextX,
			y: y,
		})
	}

	if prevX > -1 && prevY > -1 {
		coordinates = append(coordinates, coordinate{
			x: prevX,
			y: prevY,
		})
	}

	if prevX > -1 && nextY < width {
		coordinates = append(coordinates, coordinate{
			x: prevX,
			y: nextY,
		})
	}

	if prevY > -1 {
		coordinates = append(coordinates, coordinate{
			x: x,
			y: prevY,
		})
	}

	if nextY < width {
		coordinates = append(coordinates, coordinate{
			x: x,
			y: nextY,
		})
	}

	if nextX < height && prevY > -1 {
		coordinates = append(coordinates, coordinate{
			x: nextX,
			y: prevY,
		})
	}

	if nextX < height && nextY < width {
		coordinates = append(coordinates, coordinate{
			x: nextX,
			y: nextY,
		})
	}

	return coordinates
}

func neighborVisited(x, y int, path map[string]bool) (bool, *coordinate) {
	hash := getVertexHash(x, y)
	if _, found := path[hash]; found {
		return true, nil
	} else {
		c := coordinate{
			x: x,
			y: y,
		}

		return false, &c
	}
}
