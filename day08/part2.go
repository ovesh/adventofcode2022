package day08

func visibleScore(row, col int, orig [][]int) int {
	lookingRight := 0
	for i := col + 1; i < len(orig); i++ {
		lookingRight++
		if orig[row][i] >= orig[row][col] {
			break
		}
	}
	lookingLeft := 0
	for i := col - 1; i >= 0; i-- {
		lookingLeft++
		if orig[row][i] >= orig[row][col] {
			break
		}
	}
	lookingDown := 0
	for i := row + 1; i < len(orig[row]); i++ {
		lookingDown++
		if orig[i][col] >= orig[row][col] {
			break
		}
	}
	lookingUp := 0
	for i := row - 1; i >= 0; i-- {
		lookingUp++
		if orig[i][col] >= orig[row][col] {
			break
		}
	}
	return lookingRight * lookingLeft * lookingDown * lookingUp
}

func Part2() int {
	orig := parseInput(input0)
	max := 0
	for i, row := range orig {
		for j := range row {
			cur := visibleScore(i, j, orig)
			if cur > max {
				max = cur
			}
		}
	}
	return max
}
