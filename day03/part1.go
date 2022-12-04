package day03

import (
	"strings"
)

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	}
	return 0
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for _, line := range lines {
		half1 := line[:len(line)/2]
		half2 := line[len(line)/2:]
		intersection := intersect(half1, half2)
		if len(intersection) != 1 {
			panic("unexpected intersection")
		}
		var key rune
		for k := range intersection {
			key = k
		}
		total += priority(key)
	}
	return total
}
