package day01

import (
	"sort"
	"strconv"
	"strings"
)

func Part2() int {
	lines := strings.Split(input0, "\n")
	allSums := []int{}
	curTotal := 0
	for _, line := range lines {
		if line == "" {
			allSums = append(allSums, curTotal)
			curTotal = 0
			continue
		}

		curr, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		curTotal += curr
	}
	sort.Ints(allSums)
	return allSums[len(allSums)-1] + allSums[len(allSums)-2] + allSums[len(allSums)-3]
}
