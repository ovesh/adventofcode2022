package day13

import (
	"encoding/json"
	"sort"
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
	}

	var two interface{} = float64(2.0)
	var six interface{} = float64(6.0)
	allPackets = append(allPackets, []interface{}{two})
	allPackets = append(allPackets, []interface{}{six})
	sort.Slice(allPackets, func(i, j int) bool {
		return compareLists(allPackets[i], allPackets[j]) == -1
	})
	//for _, packet := range allPackets {
	//	fmt.Println(packet)
	//}
	product := 1
	for i, item := range allPackets {
		//iitem := item.([]interface{})
		if len(item) == 1 && (item[0] == two || item[0] == six) {
			product *= (i + 1)
		}
	}
	return product
}
