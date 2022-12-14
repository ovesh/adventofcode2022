package day14

func turn2(matrix *mtrx) bool {
	curX, curY := startX, startY
	for {
		if matrix[curY+1][curX] == 0 {
			// go down
			curY++
		} else if matrix[curY+1][curX-1] == 0 {
			curX--
			curY++
		} else if matrix[curY+1][curX+1] == 0 {
			curX++
			curY++
		} else {
			if curX == startX && curY == startY {
				return false
			}
			matrix[curY][curX] = 2
			return true
		}
	}
}

func Part2() int {
	matrix := prepareMatrix(input0, true)
	turns := 0
	for ; true; turns++ {
		if !turn2(&matrix) {
			break
		}
	}
	//printMatrix(&matrix)
	return turns + 1
}
