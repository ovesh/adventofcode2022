package day04

import (
	"strconv"
	"strings"
)

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getRange(s string) (int, int) {
	parts := strings.Split(s, "-")
	return mustAtoi(parts[0]), mustAtoi(parts[1])
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for _, line := range lines {
		both := strings.Split(line, ",")
		elf1Raw := both[0]
		elf2Raw := both[1]
		elf1Start, elf1End := getRange(elf1Raw)
		elf2Start, elf2End := getRange(elf2Raw)
		if elf1Start <= elf2Start && elf2End <= elf1End {
			total++
		} else if elf2Start <= elf1Start && elf1End <= elf2End {
			total++
		}
	}
	return total
}
