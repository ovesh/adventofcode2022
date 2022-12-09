package day09

import "strings"

func Part2() int {
	knotPositions := [10]pos{}
	coveredPositions := map[pos]struct{}{knotPositions[9]: {}}
	lines := strings.Split(input0, "\n")
	for _, line := range lines {
		both := strings.Split(line, " ")
		dir := both[0]
		dist := mustAtoi(both[1])
		for i := 0; i < dist; i++ {
			moveHead(&knotPositions[0], dir)
			for j := 0; j < 9; j++ {
				moveOnce(&knotPositions[j], &knotPositions[j+1], coveredPositions, j == 8)
			}
		}
	}
	return len(coveredPositions)
}
