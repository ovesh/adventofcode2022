package day02

import (
	"strings"
)

var what2play = map[string]shape{
	"A X": scissors,
	"A Y": rock,
	"A Z": paper,
	"B X": rock,
	"B Y": paper,
	"B Z": scissors,
	"C X": paper,
	"C Y": scissors,
	"C Z": rock,
}

var input2instruction = map[string]shape{
	"X": -1,
	"Y": 0,
	"Z": 1,
}

func Part2() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for _, line := range lines {
		next := what2play[line]
		both := strings.Split(line, " ")
		total += score(next, input2instruction[both[1]])
	}
	return total
}
