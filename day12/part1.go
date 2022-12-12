package day12

import (
	"strings"
)

type node struct {
	// up, right, down, left (nil for edges or if height is higher)
	neighbors [4]*node
	visited   bool
	height    int
}

func findPath(start, end *node) int {
	start.visited = true
	defer func() { start.visited = false }()
	min := 1000000
	for _, n := range start.neighbors {
		if n == nil {
			continue
		}
		if n.visited {
			continue
		}
		if n == end {
			return 1
		}
		length := findPath(n, end)
		if length == -1 { // dead end
			continue
		}
		if length < min {
			min = length + 1
		}
	}
	if min == 1000000 {
		return -1
	}
	return min
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	width := len(lines[0])
	height := len(lines)
	matrix := make([][]*node, height)
	for row := 0; row < height; row++ {
		matrix[row] = make([]*node, width)
	}
	var start *node
	var end *node
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
		}
	}
	for row, line := range lines {
		for col := range line {
			if row > 0 && matrix[row-1][col].height <= matrix[row][col].height+1 {
				matrix[row][col].neighbors[0] = matrix[row-1][col]
			}
			if row < height-1 && matrix[row+1][col].height <= matrix[row][col].height+1 {
				matrix[row][col].neighbors[2] = matrix[row+1][col]
			}
			if col > 0 && matrix[row][col-1].height <= matrix[row][col].height+1 {
				matrix[row][col].neighbors[3] = matrix[row][col-1]
			}
			if col < width-1 && matrix[row][col+1].height <= matrix[row][col].height+1 {
				matrix[row][col].neighbors[1] = matrix[row][col+1]
			}
		}
	}
	return findPath(start, end)
}
