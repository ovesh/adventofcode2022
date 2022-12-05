package day05

import (
	"strconv"
	"strings"
)

type stack struct {
	items []rune
	index int
}

func newStack(capacity int) *stack {
	return &stack{items: make([]rune, capacity)}
}

func (s *stack) push(r rune) {
	if s.index == len(s.items) {
		s.items = append(s.items, r)
	} else {
		s.items[s.index] = r
	}
	s.index++
}

func (s *stack) pop() rune {
	if s.index == 0 {
		panic("pop from empty stack")
	}
	s.index--
	return s.items[s.index]
}

func (s *stack) top() rune {
	return s.items[s.index-1]
}

func (s *stack) String() string {
	return string(s.items[0:s.index])
}

func (s *stack) len() int {
	return s.index
}

func mustAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

type instruction struct {
	from, to, count int
}

func parseInput(input string) ([]*stack, []instruction) {
	lines := strings.Split(input, "\n")
	instructions := []instruction{}
	firstInstructionIdx := -1
	for i, line := range lines {
		if strings.Index(line, "move ") == 0 {
			if firstInstructionIdx == -1 {
				firstInstructionIdx = i
			}
			iFrom := strings.Index(line, " from ")
			iTo := strings.Index(line, " to ")
			sCount := line[5:iFrom]
			sFrom := line[iFrom+6 : iTo]
			sTo := line[iTo+4:]
			instructions = append(instructions, instruction{from: mustAtoi(sFrom) - 1, to: mustAtoi(sTo) - 1, count: mustAtoi(sCount)})
		}
	}
	if firstInstructionIdx == -1 {
		panic("no instructions found")
	}

	lastStackIdx := firstInstructionIdx - 3
	numCols := (len(lines[lastStackIdx]) / 4) + 1
	stacks := make([]*stack, numCols)
	for i := 0; i < numCols; i++ {
		stacks[i] = newStack(firstInstructionIdx - 3)
	}
	for i := lastStackIdx; i >= 0; i-- {
		for colIdx := 0; colIdx < numCols; colIdx++ {
			line := lines[i]
			if len(line) < (colIdx*4)+1 {
				continue
			}
			tag := line[(colIdx*4)+1 : (colIdx*4)+2]
			if tag != " " {
				stacks[colIdx].push(rune(tag[0]))
			}
		}
	}
	return stacks, instructions
}

func Part1() string {
	stacks, instructions := parseInput(input0)
	for _, instruction := range instructions {
		for i := 0; i < instruction.count; i++ {
			popped := stacks[instruction.from].pop()
			stacks[instruction.to].push(popped)
		}
	}
	res := ""
	for _, stack := range stacks {
		if stack.len() > 0 {
			res += string(stack.pop())
		}
	}
	return res
}
