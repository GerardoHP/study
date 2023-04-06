package trees

import (
	"math"
	"study/interfaces"
)

type graphDistance struct {
	Cost int
	Path []int
}

type GraphDistances struct {
}

func (g GraphDistances) Execute() {
	//TODO implement me
	panic("implement me")
}

var _ interfaces.Exercise = (*GraphDistances)(nil)

func (GraphDistances) Solution(g [][]int, s int) []int {
	bestPath := calculate(g, s)
	var r []int
	for _, v := range bestPath {
		r = append(r, v.Cost)
	}

	return r
}

func calculate(graph [][]int, startIndex int) []graphDistance {
	visited := make(map[int]bool)                  // The visited nodes
	bestPaths := make([]graphDistance, len(graph)) // The best paths to each node from the start index
	for i := range bestPaths {
		bestPaths[i] = *newGraphDistance() // Initializing all the best paths with the maximum
	}

	bestPaths[startIndex].Cost = 0                                              // Setting the best path of the starting index to 0
	bestPaths[startIndex].Path = append(bestPaths[startIndex].Path, startIndex) // Setting the path to itself starting index

	for visitedCount := 0; visitedCount < len(graph); visitedCount++ {
		lowestCost := math.MaxInt // The lowest cost set to max integer
		lowestCostVertex := -1    // The vertex of the lowest cost

		for vertex := 0; vertex < len(graph); vertex++ {
			if _, found := visited[vertex]; !found && bestPaths[vertex].Cost < lowestCost {
				lowestCost = bestPaths[vertex].Cost
				lowestCostVertex = vertex
			}
		}

		visited[lowestCostVertex] = true
		for vertex := 0; vertex < len(graph); vertex++ {
			if _, found := visited[vertex]; !found && (graph[lowestCostVertex][vertex] != -1) {
				newPathCost := lowestCost + graph[lowestCostVertex][vertex]
				if newPathCost < bestPaths[vertex].Cost {
					bestPaths[vertex].Cost = newPathCost
					bestPaths[vertex].Path = bestPaths[lowestCostVertex].Path
					bestPaths[vertex].Path = append(bestPaths[vertex].Path, vertex)
				}
			}
		}

	}

	return bestPaths
}

func newGraphDistance() *graphDistance {
	return &graphDistance{
		Cost: math.MaxInt,
		Path: []int{},
	}
}
