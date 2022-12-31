package day17

type state struct {
	topPositions [7]int
	shapeIdx     int
	timeMod      int
}

func getTopPositions(all [][maxX + 1]bool) [maxX + 1]int {
	res := [maxX + 1]int{}
	min := -1
	for x := 0; x <= maxX; x++ {
		none := true
		for y := len(all) - 1; y >= 0; y-- {
			if all[y][x] {
				none = false
				res[x] = y + 1
				if min == -1 || y < min {
					min = y
				}
				break
			}
		}
		if none {
			min = 0
		}
	}
	for x := 0; x <= maxX; x++ {
		res[x] = res[x] - min
	}
	return res
}

func Part2() int {
	const input = input0
	const turns = 1000000000000
	knownStates := map[state]struct{}{}
	maxY := 0
	time := -1
	all := [][7]bool{}
	maxYWhenLoopDetected := -1
	turnIdxWhenLoopDetected := -1
	actualTurns := turns
	turnsBetweenLoops := 0
	maxYDiff := 0

	for turnIdx := 0; turnIdx < actualTurns; turnIdx++ {
		curShape := shapes[turnIdx%5].clone()
		curShape.leftX = 2
		curShape.bottomY = maxY + 3
		for len(all) < curShape.bottomY+curShape.height {
			all = append(all, [7]bool{})
		}

		if turnIdx > 0 {
			curState := state{getTopPositions(all), turnIdx % len(shapes), time % len(input)}
			if _, ok := knownStates[curState]; ok {
				// reset set for next loop detection
				knownStates = map[state]struct{}{}
				if maxYWhenLoopDetected != -1 {
					turnsBetweenLoops = turnIdx - turnIdxWhenLoopDetected
					maxYDiff = maxY - maxYWhenLoopDetected
					// skip to the end, we know the pattern
					actualTurns = turnIdx + ((turns - turnIdx) % turnsBetweenLoops)
				}
				turnIdxWhenLoopDetected = turnIdx
				maxYWhenLoopDetected = maxY
			}
			knownStates[curState] = struct{}{}
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
	return maxY + ((turns-actualTurns)/turnsBetweenLoops)*maxYDiff
}
