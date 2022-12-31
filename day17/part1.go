package day17

import "fmt"

type shape struct {
	shape          [][]bool
	width, height  int
	leftX, bottomY int
	rightMost      [][2]int
	leftMost       [][2]int
	bottomMost     [][2]int
}

func (s shape) clone() shape {
	return shape{s.shape, s.width, s.height, s.leftX, s.bottomY, s.rightMost, s.leftMost, s.bottomMost}
}

// all coords are bottom to top, same as the global matrix
var shapes = []shape{
	{ // ####
		shape: [][]bool{{true, true, true, true}},
		width: 4, height: 1,
		rightMost:  [][2]int{{0, 3}},
		leftMost:   [][2]int{{0, 0}},
		bottomMost: [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	}, {
		//  #
		// ###
		//  #
		shape: [][]bool{{false, true, false}, {true, true, true}, {false, true, false}},
		width: 3, height: 3,
		rightMost:  [][2]int{{0, 1}, {1, 2}, {2, 1}},
		leftMost:   [][2]int{{0, 1}, {1, 0}, {2, 1}},
		bottomMost: [][2]int{{1, 0}, {0, 1}, {1, 2}},
	}, {
		//   #
		//   #
		// ###
		shape: [][]bool{{true, true, true}, {false, false, true}, {false, false, true}},
		width: 3, height: 3,
		rightMost:  [][2]int{{0, 2}, {1, 2}, {2, 2}},
		leftMost:   [][2]int{{0, 0}, {1, 2}, {2, 2}},
		bottomMost: [][2]int{{0, 0}, {0, 1}, {0, 2}},
	}, {
		// #
		// #
		// #
		// #
		shape: [][]bool{{true}, {true}, {true}, {true}},
		width: 1, height: 4,
		rightMost:  [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		leftMost:   [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		bottomMost: [][2]int{{0, 0}},
	}, {
		// ##
		// ##
		shape: [][]bool{{true, true}, {true, true}},
		width: 2, height: 2,
		rightMost:  [][2]int{{0, 1}, {1, 1}},
		leftMost:   [][2]int{{0, 0}, {1, 0}},
		bottomMost: [][2]int{{0, 0}, {0, 1}},
	},
}

func draw(all [][7]bool) {
	for y := len(all) - 1; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < len(all[y]); x++ {
			if all[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("---------\n\n")
}

const maxX = 6

func shapeGoDown(time int, maxY int, input string, curShape *shape, all [][maxX + 1]bool) (int, bool) {
	right := input[time%len(input)] == '>'
	if right {
		canRight := curShape.leftX+curShape.width <= maxX
		for _, coords := range curShape.rightMost {
			y, x := coords[0]+curShape.bottomY, coords[1]+curShape.leftX+1
			canRight = canRight && !all[y][x]
		}
		if canRight {
			curShape.leftX++
		}
	} else {
		canLeft := curShape.leftX > 0
		for _, coords := range curShape.leftMost {
			y, x := coords[0]+curShape.bottomY, coords[1]+curShape.leftX-1
			canLeft = canLeft && !all[y][x]
		}
		if canLeft {
			curShape.leftX--
		}
	}

	if curShape.bottomY > maxY {
		curShape.bottomY--
	} else {
		canDown := true
		for _, coords := range curShape.bottomMost {
			y := curShape.bottomY + coords[0] - 1
			canDown = canDown && y >= 0 && !all[y][coords[1]+curShape.leftX]
		}
		if canDown {
			curShape.bottomY--
		} else { // can't go down any more
			for y, row := range curShape.shape {
				for x, col := range row {
					_y, _x := y+curShape.bottomY, x+curShape.leftX
					all[_y][_x] = all[_y][_x] || col
				}
			}
			nextMaxY := curShape.bottomY + curShape.height
			if nextMaxY > maxY {
				maxY = nextMaxY
			}
			//fmt.Println("maxY", maxY)
			//draw(all)
			//fmt.Println("AAAA", "time", time, "maxY", maxY)
			return maxY, false
		}
	}

	//fmt.Println("BBBB", "time", time, "maxY", maxY)
	return maxY, true
}

func Part1() int {
	const input = input0
	const turns = 2022
	maxY := 0
	time := -1
	all := [][maxX + 1]bool{}
	for i := 0; i < turns; i++ {
		curShape := shapes[i%5].clone()
		curShape.leftX = 2
		curShape.bottomY = maxY + 3
		for len(all) < curShape.bottomY+curShape.height {
			all = append(all, [maxX + 1]bool{})
		}

		for {
			time++
			var cont bool
			maxY, cont = shapeGoDown(time, maxY, input, &curShape, all)
			if !cont {
				break
			}
		}
	}
	return maxY
}
