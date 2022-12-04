package day03

import (
	"strings"
)

func groupScore(elf1, elf2, elf3 string) int {
	intersection := intersect(elf1, elf2, elf3)
	if len(intersection) != 1 {
		panic("unexpected intersection")
	}
	var key rune
	for k := range intersection {
		key = k
	}
	return priority(key)
}

func Part2() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for i := 0; i < len(lines); i += 3 {
		elf1 := lines[i]
		elf2 := lines[i+1]
		elf3 := lines[i+2]
		total += groupScore(elf1, elf2, elf3)
	}
	return total
}
