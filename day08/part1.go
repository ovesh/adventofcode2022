package day08

import (
	"strconv"
	"strings"
)

func mustAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	res := make([][]int, len(lines))
	for i, line := range lines {
		res[i] = make([]int, len(line))
		for j, c := range line {
			res[i][j] = int(c - '0')
		}
	}
	return res
}

func fromLeft(orig [][]int) [][]int {
	res := make([][]int, len(orig))
	for i, row := range orig {
		res[i] = make([]int, len(row))
		curMax := 0
		for j, v := range row {
			if v > curMax {
				curMax = v
			}
			res[i][j] = curMax
		}
	}
	return res
}

func fromRight(orig [][]int) [][]int {
	res := make([][]int, len(orig))
	for i, row := range orig {
		res[i] = make([]int, len(row))
		curMax := 0
		for j := len(row) - 1; j >= 0; j-- {
			v := row[j]
			if v > curMax {
				curMax = v
			}
			res[i][j] = curMax
		}
	}
	return res
}

func fromTop(orig [][]int) [][]int {
	res := make([][]int, len(orig))
	for i, row := range orig {
		res[i] = make([]int, len(row))
	}
	for i := 0; i < len(orig); i++ {
		curMax := 0
		for j := 0; j < len(orig[i]); j++ {
			v := orig[j][i]
			if v > curMax {
				curMax = v
			}
			res[j][i] = curMax
		}
	}
	return res
}

func fromBottom(orig [][]int) [][]int {
	res := make([][]int, len(orig))
	for i, row := range orig {
		res[i] = make([]int, len(row))
	}
	for i := 0; i < len(orig); i++ {
		curMax := 0
		for j := len(orig) - 1; j >= 0; j-- {
			v := orig[j][i]
			if v > curMax {
				curMax = v
			}
			res[j][i] = curMax
		}
	}
	return res
}

func isVisibleFromLeft(col, row int, orig [][]int, viewFromLeft [][]int) bool {
	return orig[row][col] <= viewFromLeft[row][col] && viewFromLeft[row][col-1] != viewFromLeft[row][col]
}

func isVisibleFromRight(col, row int, orig [][]int, viewFromRight [][]int) bool {
	return orig[row][col] <= viewFromRight[row][col] && viewFromRight[row][col+1] != viewFromRight[row][col]
}

func isVisibleFromTop(col, row int, orig [][]int, viewFromTop [][]int) bool {
	return orig[row][col] <= viewFromTop[row][col] && viewFromTop[row-1][col] != viewFromTop[row][col]
}

func isVisibleFromBottom(col, row int, orig [][]int, viewFromBottom [][]int) bool {
	return orig[row][col] <= viewFromBottom[row][col] && viewFromBottom[row+1][col] != viewFromBottom[row][col]
}

func Part1() int {
	orig := parseInput(input0)
	left := fromLeft(orig)
	top := fromTop(orig)
	right := fromRight(orig)
	bottom := fromBottom(orig)
	total := 0
	for i := 1; i < len(left)-1; i++ {
		for j := 1; j < len(left[i])-1; j++ {
			visibleFromLeft := isVisibleFromLeft(j, i, orig, left)
			visibleFromTop := isVisibleFromTop(j, i, orig, top)
			visibleFromRight := isVisibleFromRight(j, i, orig, right)
			visibleFromBottom := isVisibleFromBottom(j, i, orig, bottom)
			if visibleFromLeft || visibleFromTop || visibleFromRight || visibleFromBottom {
				//fmt.Println("adding i=", i, "j=", j, "visibleFromLeft=", visibleFromLeft, "visibleFromTop=", visibleFromTop, "visibleFromRight=", visibleFromRight, "visibleFromBottom=", visibleFromBottom)
				total++
			}
		}
	}
	return total + 4*len(orig) - 4
}
