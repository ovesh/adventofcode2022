package day03

import (
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
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
		half1 := sortString(line[:len(line)/2])
		half2 := sortString(line[len(line)/2:])
		idx1 := 0
		idx2 := 0
		last := uint8(0)
		for idx1 < len(half1) && idx2 < len(half2) {
			if half1[idx1] == half2[idx2] {
				if half1[idx1] != last {
					last = half1[idx1]
					total += priority(rune(last))
				}
				idx1++
				idx2++
			} else if half1[idx1] < half2[idx2] {
				idx1++
			} else {
				idx2++
			}
		}
	}
	return total
}
