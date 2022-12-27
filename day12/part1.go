package day12

import (
	"fmt"
	"strings"

	structs "github.com/kirves/godijkstra/common/structs"
	"github.com/kirves/godijkstra/dijkstra"
)

type node struct {
	height int
	row    int
	col    int
}

func (n *node) String() string {
	return fmt.Sprintf("(%d,%d),", n.row, n.col)
}

func (g *GraphObject) SuccessorsForNode(node string) []structs.Connection {
	return g.nodes2Neighbors[node]
}
func (g *GraphObject) PredecessorsFromNode(_ string) []structs.Connection {
	return []structs.Connection{}
}

func (g *GraphObject) EdgeWeight(_, _ string) int {
	return 1
}

type GraphObject struct {
	nodes2Neighbors map[string][]structs.Connection
}

func setup(input string) (*GraphObject, *node, *node, [][]*node) {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)
	matrix := make([][]*node, height)
	for row := 0; row < height; row++ {
		matrix[row] = make([]*node, width)
	}
	var start *node
	var end *node
	graph := GraphObject{nodes2Neighbors: map[string][]structs.Connection{}}
	for row, line := range lines {
		for col, c := range line {
			if c == 'S' {
				matrix[row][col] = &node{height: 0}
				start = matrix[row][col]
			} else if c == 'E' {
				matrix[row][col] = &node{height: 'z' - 'a'}
				end = matrix[row][col]
			} else {
				matrix[row][col] = &node{height: int(c - 'a')}
			}
			matrix[row][col].row = row
			matrix[row][col].col = col
			graph.nodes2Neighbors[matrix[row][col].String()] = []structs.Connection{}
		}
	}
	for row, line := range lines {
		for col := range line {
			nodeID := matrix[row][col].String()
			if row > 0 && matrix[row-1][col].height <= matrix[row][col].height+1 {
				graph.nodes2Neighbors[nodeID] = append(graph.nodes2Neighbors[nodeID], structs.Connection{Destination: (&node{row: row - 1, col: col}).String(), Weight: 1})
			}
			if row < height-1 && matrix[row+1][col].height <= matrix[row][col].height+1 {
				graph.nodes2Neighbors[nodeID] = append(graph.nodes2Neighbors[nodeID], structs.Connection{Destination: (&node{row: row + 1, col: col}).String(), Weight: 1})
			}
			if col > 0 && matrix[row][col-1].height <= matrix[row][col].height+1 {
				graph.nodes2Neighbors[nodeID] = append(graph.nodes2Neighbors[nodeID], structs.Connection{Destination: (&node{row: row, col: col - 1}).String(), Weight: 1})
			}
			if col < width-1 && matrix[row][col+1].height <= matrix[row][col].height+1 {
				graph.nodes2Neighbors[nodeID] = append(graph.nodes2Neighbors[nodeID], structs.Connection{Destination: (&node{row: row, col: col + 1}).String(), Weight: 1})
			}
		}
	}

	return &graph, start, end, matrix
}

func Part1() int {
	graph, start, end, _ := setup(input0)
	path, valid := dijkstra.SearchPath(graph, start.String(), end.String(), dijkstra.VANILLA)
	if !valid {
		panic("invalid path")
	}
	fmt.Println(path.Path)
	return len(path.Path) - 1
}
