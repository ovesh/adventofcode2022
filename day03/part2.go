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
