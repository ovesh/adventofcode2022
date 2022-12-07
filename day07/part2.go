package day07

func Part2() int {
	root, allNodes := parseInput(input0)
	remainingSize := 70000000 - root.size
	curMin := root.size
	min2Delete := 30000000 - remainingSize
	for _, n := range allNodes {
		if n.size < min2Delete {
			continue
		}
		if n.size < curMin {
			curMin = n.size
		}
	}
	return curMin
}
