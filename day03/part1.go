package day03

import (
	"strings"
)

type runeset map[rune]struct{}

func (r runeset) addAll(s string) {
	for _, c := range s {
		r[c] = struct{}{}
	}
}

func (r runeset) contains(c rune) bool {
	_, ok := r[c]
	return ok
}

func intersect(s ...string) runeset {
	result := runeset{}
	sets := make([]runeset, len(s))
	for i, s := range s {
		sets[i] = runeset{}
		sets[i].addAll(s)
	}

	for r := range sets[0] {
		all := true
		for i := 1; i < len(sets); i++ {
			if !sets[i].contains(r) {
				all = false
				break
			}
		}
		if all {
			result[r] = struct{}{}
		}
	}
	return result
}

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
