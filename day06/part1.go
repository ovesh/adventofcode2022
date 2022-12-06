package day06

func Part1() int {
	input := input0
	for i := 3; i < len(input); i++ {
		set := make(map[uint8]struct{})
		set[input[i-3]] = struct{}{}
		set[input[i-2]] = struct{}{}
		set[input[i-1]] = struct{}{}
		set[input[i]] = struct{}{}
		if len(set) == 4 {
			return i + 1
		}
	}
	panic("no solution")
}
