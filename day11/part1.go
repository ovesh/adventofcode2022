package day11

import (
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	operation   func(int) int
	test        func(int) bool
	targetTrue  int
	targetFalse int
	inspected   int

	// part 2 only
	denominator int
}

func (m *monkey) turn(allMonkeys []*monkey) {
	for i := range m.items {
		m.items[i] = m.operation(m.items[i])
		m.items[i] = m.items[i] / 3
		targetIdx := m.targetFalse
		if m.test(m.items[i]) {
			targetIdx = m.targetTrue
		}
		allMonkeys[targetIdx].items = append(allMonkeys[targetIdx].items, m.items[i])
	}
	m.inspected += len(m.items)
	m.items = []int{}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func sAtoi(s []string) []int {
	is := make([]int, len(s))
	for i, v := range s {
		is[i] = atoi(v)
	}
	return is
}

func setupMonkeys(s string) []*monkey {
	lines := strings.Split(s, "\n")
	monkeys := []*monkey{}
	for i := 0; i < len(lines); i += 7 {
		m := &monkey{}
		sItems := strings.Split(lines[i+1][len("  Starting items: "):], ", ")
		m.items = sAtoi(sItems)

		prefixLen := len("  Operation: new = old ")
		sOp := lines[i+2][prefixLen : prefixLen+1]
		sOperand := lines[i+2][prefixLen+2:]
		if sOperand == "old" && sOp == "*" {
			m.operation = func(i int) int { return i * i }
		} else {
			operand := atoi(sOperand)
			if sOp == "*" {
				m.operation = func(i int) int { return i * operand }
			} else { // +
				m.operation = func(i int) int { return i + operand }
			}
		}

		denominator := atoi(lines[i+3][len("  Test: divisible by "):])
		m.denominator = denominator
		m.test = func(i int) bool { return i%denominator == 0 }

		m.targetTrue = atoi(lines[i+4][len("    If true: throw to monkey "):])
		m.targetFalse = atoi(lines[i+5][len("    If false: throw to monkey "):])

		monkeys = append(monkeys, m)
	}
	return monkeys
}

func Part1() int {
	monkeys := setupMonkeys(input0)
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.turn(monkeys)
		}
		//for i, m := range monkeys {
		//	fmt.Println(i, m.inspected, m.items)
		//}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	return monkeys[0].inspected * monkeys[1].inspected
}
