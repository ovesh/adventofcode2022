package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func atoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

type mtrx [200][110]int

func printMatrix(matrix *mtrx) {
	fmt.Print("  ")
	for i := 0; i < 10; i++ {
		fmt.Print("0123456789")
	}
	fmt.Println()
	for i, row := range matrix {
		fmt.Print((i / 10) % 10)
		fmt.Print(i % 10)
		for _, cell := range row {
			if cell == 1 {
				print("#")
			} else if cell == 2 {
				print("o")
			} else {
				print(".")
			}
		}
		println()
	}
}

const xOffset = 480

func prepareMatrix(input string) mtrx {
	// matrix starts from 480, 0
	matrix := mtrx{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sPoints := strings.Split(line, " -> ")
		prevX, prevY := 0, 0
		for j, sPoint := range sPoints {
			xCoord := strings.Split(sPoint, ",")
			x := atoi(xCoord[0])
			y := atoi(xCoord[1])
			if j == 0 {
				prevX = x
				prevY = y
				continue
			}
			if x > prevX { // y == prevY
				for i := prevX; i <= x; i++ {
					matrix[y][i-xOffset] = 1
				}
			}
			if x < prevX { // y == prevY
				for i := x; i <= prevX; i++ {
					matrix[y][i-xOffset] = 1
				}
			}
			if y > prevY { // x == prevX
				for i := prevY; i <= y; i++ {
					matrix[i][x-xOffset] = 1
				}
			}
			if y < prevY { // x == prevX
				for i := y; i <= prevY; i++ {
					matrix[i][x-xOffset] = 1
				}
			}
			prevX, prevY = x, y
		}
	}
	return matrix
}

const startX, startY = 500 - xOffset, 0

func turn(matrix *mtrx) bool {
	curX, curY := startX, startY
	for {
		if curY == 199 {
			return false
		}
		if matrix[curY+1][curX] == 0 {
			// go down
			curY++
		} else if matrix[curY+1][curX-1] == 0 {
			curX--
			curY++
		} else if matrix[curY+1][curX+1] == 0 {
			curX++
			curY++
		} else {
			matrix[curY][curX] = 2
			return true
		}
	}
}

func Part1() int {
	matrix := prepareMatrix(input0)
	turns := 0
	for ; true; turns++ {
		if !turn(&matrix) {
			break
		}
	}
	printMatrix(&matrix)
	return turns
}
