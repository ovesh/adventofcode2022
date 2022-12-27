package day12

import (
	"github.com/kirves/godijkstra/dijkstra"
)

func Part2() int {
	graph, _, end, matrix := setup(input0)
	starts := []*node{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col].height == 0 {
				starts = append(starts, matrix[row][col])
			}
		}
	}
	min := 1000000000
	for _, start := range starts {
		path, valid := dijkstra.SearchPath(graph, start.String(), end.String(), dijkstra.VANILLA)
		if !valid {
			continue
		}
		if len(path.Path) < min {
			min = len(path.Path)
		}
	}
	return min - 1
}
