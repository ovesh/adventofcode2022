package day10

import (
	"fmt"
	"strings"
)

func printCRT(crt [6][40]rune) {
	for i := range crt {
		for j := range crt[i] {
			fmt.Printf("%c", crt[i][j])
		}
		fmt.Println()
	}
}

func updateCRT(cycle int, register int, crt *[6][40]rune) {
	row := ((cycle - 1) / 40) % 6
	col := (cycle - 1) % 40
	fmt.Println(cycle, row, col)
	if col == register || col == register+1 || col == register-1 {
		crt[row][col] = '#'
	} else {
		crt[row][col] = '.'
	}
}

func Part2() int {
	crt := [6][40]rune{}
	for i := range crt {
		for j := range crt[i] {
			crt[i][j] = ' '
		}
	}
	lines := strings.Split(input0, "\n")
	cycle := 1
	register := 1
	for _, line := range lines {
		if line == "noop" {
			updateCRT(cycle, register, &crt)
			cycle++
		} else if strings.HasPrefix(line, "addx ") {
			updateCRT(cycle, register, &crt)
			cycle++
			//fmt.Println(cycle)
			updateCRT(cycle, register, &crt)
			register += mustAtoi(line[5:])
			cycle++
		}
	}
	printCRT(crt)
	return 0
}
