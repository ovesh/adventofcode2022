package day13

import (
	"encoding/json"
	"strings"
)

// return -1 for equal, 0 for "right order", 1 for "wrong order"
func compareLists(list1, list2 []interface{}) int {
	for i := 0; i < len(list1) || i < len(list2); i++ {
		if i >= len(list1) {
			return 0
		}
		if i >= len(list2) {
			return 1
		}
		_, isInt1 := list1[i].(float64) // ugh, json unmarshals numbers as float64
		_, isInt2 := list2[i].(float64)
		if isInt1 && isInt2 {
			if int(list1[i].(float64)) == int(list2[i].(float64)) {
				continue
			}
			if int(list1[i].(float64)) < int(list2[i].(float64)) {
				return 0
			}
			return 1
		}
		if isInt1 && !isInt2 {
			winner := compareLists(([]interface{}{list1[i]}), list2[i].([]interface{}))
			if winner == -1 { // draw
				continue
			}
			return winner
		}
		if !isInt1 && isInt2 {
			winner := compareLists(list1[i].([]interface{}), ([]interface{}{list2[i]}))
			if winner == -1 { // draw
				continue
			}
			return winner
		}
		winner := compareLists(list1[i].([]interface{}), list2[i].([]interface{}))
		if winner == -1 { // draw
			continue
		}
		return winner
	}

	return -1
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	total := 0
	for i := 0; i < len(lines); i += 3 {
		sList1 := lines[i]
		sList2 := lines[i+1]
		var list1, list2 []interface{}
		if err := json.Unmarshal([]byte(sList1), &list1); err != nil {
			panic(err)
		}
		if err := json.Unmarshal([]byte(sList2), &list2); err != nil {
			panic(err)
		}
		//fmt.Println(lines[i], "vs", lines[i+1])
		winner := compareLists(list1, list2)
		if winner == -1 {
			panic("draw")
		}
		if winner == 0 {
			total += (i / 3) + 1
		}
	}
	return total
}
