package day04

import "strings"

func Part2() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for _, line := range lines {
		both := strings.Split(line, ",")
		elf1Raw := both[0]
		elf2Raw := both[1]
		elf1Start, elf1End := getRange(elf1Raw)
		elf2Start, elf2End := getRange(elf2Raw)
		if !(elf1End < elf2Start || elf2End < elf1Start) {
			total++
		}
	}
	return total
}
