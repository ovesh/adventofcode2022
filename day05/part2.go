package day05

func (s *stack) pushS(str string) {
	for _, r := range str {
		s.push(r)
	}
}

func (s *stack) popS(len int) string {
	res := make([]rune, len)
	for i := len - 1; i >= 0; i-- {
		res[i] = s.pop()
	}
	return string(res)
}

func Part2() string {
	stacks, instructions := parseInput(input0)
	for _, instruction := range instructions {
		popped := stacks[instruction.from].popS(instruction.count)
		stacks[instruction.to].pushS(popped)
	}
	res := ""
	for _, stack := range stacks {
		if stack.len() > 0 {
			res += string(stack.pop())
		}
	}
	return res
}
