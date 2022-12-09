package day09

import (
	"strconv"
	"strings"
)

type pos struct{ x, y int }

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func moveHead(posH *pos, dir string) {
	switch dir {
	case "R":
		posH.x++
	case "L":
		posH.x--
	case "U":
		posH.y++
	case "D":
		posH.y--
	}
}

func moveOnce(posH, posT *pos, positions map[pos]struct{}, update bool) {
	defer func() {
		if update {
			positions[*posT] = struct{}{}
		}
	}()
	if *posH == *posT {
		return
	}
	if posH.x == posT.x {
		if abs(posH.y-posT.y) == 1 {
			return
		}
		if posH.y > posT.y {
			posT.y++
		} else {
			posT.y--
		}
		return
	}
	if posH.y == posT.y {
		if abs(posH.x-posT.x) == 1 {
			return
		}
		if posH.x > posT.x {
			posT.x++
		} else {
			posT.x--
		}
		return
	}
	if abs(posH.x-posT.x) == 1 && abs(posH.y-posT.y) == 1 {
		return
	}
	if posH.x-posT.x == 2 {
		posT.x++
		if posH.y > posT.y {
			posT.y++
		} else {
			posT.y--
		}
		return
	}
	if posH.x-posT.x == -2 {
		posT.x--
		if posH.y > posT.y {
			posT.y++
		} else {
			posT.y--
		}
		return
	}
	if posH.y-posT.y == 2 {
		posT.y++
		if posH.x > posT.x {
			posT.x++
		} else {
			posT.x--
		}
		return
	}
	if posH.y-posT.y == -2 {
		posT.y--
		if posH.x > posT.x {
			posT.x++
		} else {
			posT.x--
		}
		return
	}
}

func Part1() int {
	posH := pos{0, 0}
	posT := pos{0, 0}
	coveredPositions := map[pos]struct{}{posT: {}}
	lines := strings.Split(input0, "\n")
	for _, line := range lines {
		both := strings.Split(line, " ")
		dir := both[0]
		dist := mustAtoi(both[1])
		for i := 0; i < dist; i++ {
			moveHead(&posH, dir)
			moveOnce(&posH, &posT, coveredPositions, true)
		}
	}
	return len(coveredPositions)
}
