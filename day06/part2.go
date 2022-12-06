package day06

func Part2() int {
	input := input0
	for i := 13; i < len(input); i++ {
		set := make(map[uint8]struct{})
		for j := 0; j <= 13; j++ {
			set[input[i-j]] = struct{}{}
		}
		if len(set) == 14 {
			return i + 1
		}
	}
	panic("no solution")
}
