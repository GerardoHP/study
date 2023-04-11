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
	var result []string
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
				var m []coordinate
				if solveWordBoggle(board, &m, word, "", c.x, c.y) {
					result = append(result, word)
					break
				}
			}
		}
	}

	return result
}

func solveWordBoggle(board [][]string, visited *[]coordinate, word, currentWord string, x, y int) bool {
	//isPrefix := strings.HasPrefix(word, currentWord)
	//if !isPrefix && len(currentWord) > 0 {
	//	return false
	//}

	if word == currentWord {
		return true
	}

	if len(word) == len(currentWord) {
		return false
	}

	*visited = append(*visited, coordinate{x: x, y: y})
	for i := x; i < len(board); i++ {
		for j := y; j < len(board[i]); j++ {
			neighbors := getNeighbors(i, j, len(board[i]), len(board), *visited)
			newCurrentWord := fmt.Sprintf("%v%v", currentWord, board[i][j])
			for _, neighbor := range neighbors {
				if solveWordBoggle(board, visited, word, newCurrentWord, neighbor.x, neighbor.y) {
					return true
				}
			}
		}
	}

	return false
}

func getNeighbors(x, y, width, height int, visits []coordinate) []coordinate {
	var coordinates []coordinate
	nextX := x + 1
	prevX := x - 1
	nextY := y + 1
	prevY := y - 1

	if visited, c := neighborVisited(prevX, y, visits); !visited && prevX > -1 {
		coordinates = append(coordinates, *c)
	}

	if visited, c := neighborVisited(nextX, y, visits); !visited && nextX < height {
		coordinates = append(coordinates, *c)
	}

	if visited, c := neighborVisited(prevX, prevY, visits); !visited && prevX > -1 && prevY > -1 {
		coordinates = append(coordinates, *c)
	}

	if visited, c := neighborVisited(prevX, nextY, visits); !visited && prevX > -1 && nextY < width {
		coordinates = append(coordinates, *c)
	}

	if visited, c := neighborVisited(x, prevY, visits); !visited && prevY > -1 {
		coordinates = append(coordinates, *c)
	}

	if visited, c := neighborVisited(x, nextY, visits); !visited && nextY < width {
		coordinates = append(coordinates, *c)
	}

	if visited, c := neighborVisited(nextX, prevY, visits); !visited && nextX < height && prevY > -1 {
		coordinates = append(coordinates, *c)
	}

	if visited, c := neighborVisited(nextX, nextY, visits); !visited && nextX < height && nextY < width {
		coordinates = append(coordinates, *c)
	}

	return coordinates
}

func neighborVisited(x, y int, visited []coordinate) (bool, *coordinate) {
	for _, v := range visited {
		if v.x == x && v.y == y {
			return true, nil
		}
	}

	c := coordinate{
		x: x,
		y: y,
	}

	return false, &c
}
