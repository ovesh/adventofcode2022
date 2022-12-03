package day03

import (
	"strings"
)

func min(r1, r2, r3 uint8) uint8 {
	if r1 < r2 && r1 < r3 {
		return r1
	}
	if r2 < r3 {
		return r2
	}
	return r3
}

func groupScore(elf1, elf2, elf3 string) int {
	idx1 := 0
	idx2 := 0
	idx3 := 0
	for idx1 < len(elf1) && idx2 < len(elf2) && idx3 < len(elf3) {
		cur1 := elf1[idx1]
		cur2 := elf2[idx2]
		cur3 := elf3[idx3]
		if cur1 == cur2 && cur2 == cur3 {
			return priority(rune(cur1))
		}
		minRune := min(cur1, cur2, cur3)
		if minRune == cur1 {
			idx1++
		}
		if minRune == cur2 {
			idx2++
		}
		if minRune == cur3 {
			idx3++
		}
	}
	panic("no match")
}

func Part2() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for i := 0; i < len(lines); i += 3 {
		elf1 := sortString(lines[i])
		elf2 := sortString(lines[i+1])
		elf3 := sortString(lines[i+2])
		total += groupScore(elf1, elf2, elf3)
	}
	return total
}
