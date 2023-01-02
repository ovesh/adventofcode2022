package day18

import (
	"strings"
)

type coordinate struct {
	x, y, z    int
	hasDroplet bool
	external   bool
	seen       bool
	neighbours []*coordinate
}

func markAllExternal(c *coordinate) {
	for _, coord := range c.neighbours {
		if coord.seen {
			continue
		}
		coord.seen = true
		if coord.hasDroplet {
			coord.external = true
			continue
		}
		coord.external = true
		markAllExternal(coord)
	}
}

func Part2() int {
	const input = input0

	lines := strings.Split(input, "\n")
	all := [side][side][side]*coordinate{}
	droplets := []droplet{}
	for _, line := range lines {
		sVals := strings.Split(line, ",")
		x := atoi(sVals[0])
		y := atoi(sVals[1])
		z := atoi(sVals[2])
		all[x][y][z] = &coordinate{x: x, y: y, z: z, hasDroplet: true, neighbours: []*coordinate{}}
		droplets = append(droplets, droplet{x: x, y: y, z: z})
	}

	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			for z := 0; z < side; z++ {
				if all[x][y][z] == nil {
					all[x][y][z] = &coordinate{x: x, y: y, z: z, hasDroplet: false, neighbours: []*coordinate{}}
				}
			}
		}
	}
	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			for z := 0; z < side; z++ {
				if x > 0 {
					all[x][y][z].neighbours = append(all[x][y][z].neighbours, all[x-1][y][z])
				}
				if x < side-1 {
					all[x][y][z].neighbours = append(all[x][y][z].neighbours, all[x+1][y][z])
				}
				if y > 0 {
					all[x][y][z].neighbours = append(all[x][y][z].neighbours, all[x][y-1][z])
				}
				if y < side-1 {
					all[x][y][z].neighbours = append(all[x][y][z].neighbours, all[x][y+1][z])
				}
				if z > 0 {
					all[x][y][z].neighbours = append(all[x][y][z].neighbours, all[x][y][z-1])
				}
				if z < side-1 {
					all[x][y][z].neighbours = append(all[x][y][z].neighbours, all[x][y][z+1])
				}
			}
		}
	}

	all[0][0][0].external = true // the input data doesn't list 0,0,0
	markAllExternal(all[0][0][0])
	total := 0
	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			for z := 0; z < side; z++ {
				coord := all[x][y][z]
				if !coord.external {
					continue
				}
				if !coord.hasDroplet {
					continue
				}
				for _, neighbour := range coord.neighbours {
					if !neighbour.hasDroplet && neighbour.external {
						total++
					}
				}
				if len(coord.neighbours) != 6 {
					total += (6 - len(coord.neighbours))
				}
			}
		}
	}

	return total
}
