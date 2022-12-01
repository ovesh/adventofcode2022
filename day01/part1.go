package day01

import (
	"strconv"
	"strings"
)

func Part1() int {
	lines := strings.Split(input0, "\n")
	max := 0
	curTotal := 0
	for _, line := range lines {
		if line == "" {
			if curTotal > max {
				max = curTotal
			}
			curTotal = 0
			continue
		}

		curr, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		curTotal += curr
	}

	return max
}
