package backtracking

import (
	"fmt"
	"strings"
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

	initial := make(map[uint8][]coordinate)
	for x, v := range board {
		for y, v1 := range v {
			r := v1[0]
			initial[r] = append(initial[r], coordinate{x, y})
		}
	}

	for _, word := range words {
		if coordinates, found := initial[word[0]]; found {
			for _, c := range coordinates {
				if solveWordBoggle(board, make(map[string]bool), word, "", c.x, c.y) {
					result = append(result, word)
					break
				}
			}
		}

		//if solveWordBoggle(board, make(map[string]bool), word, "", 0, 0) {
		//	result = append(result, word)
		//}
	}

	return result
}

func solveWordBoggle(board [][]string, visited map[string]bool, word, currentWord string, x, y int) bool {
	isPrefix := strings.HasPrefix(word, currentWord)
	if !isPrefix && len(currentWord) > 0 {
		return false
	}

	if word == currentWord {
		return true
	}

	if len(word) == len(currentWord) {
		return false
	}

	//currentHash := getVertexHash(x, y)
	//visited[currentHash] = true
	for i := x; i < len(board); i++ {
		for j := y; j < len(board[i]); j++ {
			currentHash := getVertexHash(i, j)
			visited[currentHash] = true
			neighbors := getNeighbors(i, j, len(board[i]), len(board), visited)
			newCurrentWord := fmt.Sprintf("%v%v", currentWord, board[i][j])
			for _, neighbor := range neighbors {
				neighborHash := getVertexHash(i, j)
				visited[neighborHash] = true
				if solveWordBoggle(board, visited, word, newCurrentWord, neighbor.x, neighbor.y) {
					return true
				}
			}
			//if x == 0 && y == 0 {
			//	visited = make(map[string]bool)
			//}
		}
		//
		//if x == 0 && y == 0 {
		//	visited = make(map[string]bool)
		//}
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

	if exist, c := neighborVisited(prevX, y, path); !exist && prevX > -1 {
		coordinates = append(coordinates, *c)
	}

	if exist, c := neighborVisited(nextX, y, path); !exist && nextX < height {
		coordinates = append(coordinates, *c)
	}

	if exist, c := neighborVisited(prevX, prevY, path); !exist && prevX > -1 && prevY > -1 {
		coordinates = append(coordinates, *c)
	}

	if exist, c := neighborVisited(prevX, nextY, path); !exist && prevX > -1 && nextY < width {
		coordinates = append(coordinates, *c)
	}

	if exist, c := neighborVisited(x, prevY, path); !exist && prevY > -1 {
		coordinates = append(coordinates, *c)
	}

	if exist, c := neighborVisited(x, nextY, path); !exist && nextY < width {
		coordinates = append(coordinates, *c)
	}

	if exist, c := neighborVisited(nextX, prevY, path); !exist && nextX < height && prevY > -1 {
		coordinates = append(coordinates, *c)
	}

	if exist, c := neighborVisited(nextX, nextY, path); !exist && nextX < height && nextY < width {
		coordinates = append(coordinates, *c)
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
