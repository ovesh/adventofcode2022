package day11

import (
	"fmt"
	"sort"
)

func (m *monkey) turn2(allMonkeys []*monkey, mainLCM int) {
	for i := range m.items {
		m.items[i] = m.operation(m.items[i]) % mainLCM
		targetIdx := m.targetFalse
		if m.test(m.items[i]) {
			targetIdx = m.targetTrue
		}
		allMonkeys[targetIdx].items = append(allMonkeys[targetIdx].items, m.items[i])
	}
	m.inspected += len(m.items)
	m.items = []int{}
}

func max(nums ...int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func lcm(nums ...int) int {
	res := max(nums...)
	for ; true; res++ {
		allTrue := true
		for _, n := range nums {
			if res%n != 0 {
				allTrue = false
				break
			}
		}
		if allTrue {
			return res
		}
	}
	panic("unreachable")
}

func Part2() int {
	monkeys := setupMonkeys(input0)
	allDenominators := []int{}
	for _, m := range monkeys {
		allDenominators = append(allDenominators, m.denominator)
	}
	mainLCM := lcm(allDenominators...)
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.turn2(monkeys, mainLCM)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	fmt.Println(monkeys[0].inspected, monkeys[1].inspected)
	return monkeys[0].inspected * monkeys[1].inspected
}
