package day18

import (
	"strconv"
	"strings"
)

func atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

type droplet struct {
	x, y, z int
}

const side = 21

func total1(droplets []droplet, all [side][side][side]bool) int {
	total := 0
	for _, droplet := range droplets {
		if droplet.x == 0 || droplet.x == side-1 {
			total++
		}
		if droplet.y == 0 || droplet.y == side-1 {
			total++
		}
		if droplet.z == 0 || droplet.z == side-1 {
			total++
		}
		if droplet.x > 0 && !all[droplet.x-1][droplet.y][droplet.z] {
			total++
		}
		if droplet.x < side-1 && !all[droplet.x+1][droplet.y][droplet.z] {
			total++
		}
		if droplet.y > 0 && !all[droplet.x][droplet.y-1][droplet.z] {
			total++
		}
		if droplet.y < side-1 && !all[droplet.x][droplet.y+1][droplet.z] {
			total++
		}
		if droplet.z > 0 && !all[droplet.x][droplet.y][droplet.z-1] {
			total++
		}
		if droplet.z < side-1 && !all[droplet.x][droplet.y][droplet.z+1] {
			total++
		}
	}

	return total
}

func parseInput(input string) ([]droplet, [side][side][side]bool) {
	lines := strings.Split(input, "\n")
	all := [side][side][side]bool{}
	droplets := []droplet{}
	for _, line := range lines {
		sVals := strings.Split(line, ",")
		x := atoi(sVals[0])
		y := atoi(sVals[1])
		z := atoi(sVals[2])
		all[x][y][z] = true
		droplets = append(droplets, droplet{x: x, y: y, z: z})
	}
	return droplets, all
}

func Part1() int {
	const input = input0

	droplets, all := parseInput(input)
	return total1(droplets, all)
}
