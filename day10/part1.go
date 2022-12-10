package day10

import (
	"strconv"
	"strings"
)

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func update(cycle, register, curTotal int) int {
	if (cycle-20)%40 == 0 {
		//fmt.Printf("cycle %d: register %d, updateing total by %d from %d to %d\n", cycle, register, cycle*register, curTotal, curTotal+(cycle*register))
		return curTotal + (cycle * register)
	}
	return curTotal
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	cycle := 1
	register := 1
	total := 0
	for _, line := range lines {
		//fmt.Println(cycle, line)
		if line == "noop" {
			total = update(cycle, register, total)
			cycle++
		} else if strings.HasPrefix(line, "addx ") {
			total = update(cycle, register, total)
			cycle++
			//fmt.Println(cycle)
			total = update(cycle, register, total)
			register += mustAtoi(line[5:])
			cycle++
		}
	}
	return total
}
