package day13

import (
	"encoding/json"
	"strings"
)

func Part2() int {
	lines := strings.Split(input0, "\n")
	allPackets := [][]interface{}{}
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
		allPackets = append(allPackets, list1)
		allPackets = append(allPackets, list2)
		winner := compareLists(list1, list2)
		if winner == -1 {
			panic("draw")
		}
		//if winner == 0 {
		//	total += (i / 3) + 1
		//}
	}
	//return total
	return 0
}
