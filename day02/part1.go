package day02

import (
	"strings"
)

type shape = int

const (
	nope shape = iota
	rock
	paper
	scissors
)

func score(s shape, won int) int {
	res := s
	if won > 0 {
		res += 6
	} else if won == 0 {
		res += 3
	}
	return res
}

var input2Shape = map[string]shape{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

func won(s1, s2 shape) int {
	if s1 == s2 {
		return 0
	}
	if s1 == rock && s2 == scissors {
		return 1
	}
	if s1 == paper && s2 == rock {
		return 1
	}
	if s1 == scissors && s2 == paper {
		return 1
	}
	return -1
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for _, line := range lines {
		both := strings.Split(line, " ")
		them := input2Shape[both[0]]
		me := input2Shape[both[1]]
		total += score(me, won(me, them))
	}
	return total
}
