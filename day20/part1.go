package day20

import (
	"fmt"
	"strconv"
	"strings"
)

type item struct {
	val  int
	next *item
	prev *item
}

func (i *item) String() string {
	return fmt.Sprintf("%d", i.val)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func printAll(i *item) {
	first := i
	for curItem := i; curItem.next != first; curItem = curItem.next {
		fmt.Print(curItem.val, " ")
	}
	fmt.Println(i.prev.val)
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	numItems := len(lines)
	origOrder := make([]*item, numItems)
	var zero *item = nil
	for i, line := range lines {
		origOrder[i] = &item{val: atoi(line)}
		if origOrder[i].val == 0 {
			zero = origOrder[i]
		}
		if i > 0 {
			origOrder[i-1].next = origOrder[i]
			origOrder[i].prev = origOrder[i-1]
		}
	}
	//first := origOrder[0]
	origOrder[numItems-1].next = origOrder[0]
	origOrder[0].prev = origOrder[numItems-1]
	//fmt.Println(origOrder)
	for i := 0; i < numItems; i++ {
		curItem := origOrder[i]
		if curItem.val == 0 {
			fmt.Println("item", i, curItem)
			//printAll(origOrder[0])
			continue
		}
		newNext := curItem
		if curItem.val > 0 {
			for j := 0; j < curItem.val; j++ {
				newNext = newNext.next
			}
		}
		if curItem.val < 0 {
			for j := 0; j <= -curItem.val; j++ {
				newNext = newNext.prev
			}
		}
		curItem.prev.next = curItem.next
		curItem.next.prev = curItem.prev
		newNext.next.prev = curItem
		curItem.next = newNext.next
		curItem.prev = newNext
		newNext.next = curItem

		fmt.Println("item", i, curItem)
		//printAll(first)
	}
	//fmt.Println("done")
	//printAll(first)
	//fmt.Println("first", first)

	cur := zero
	res := 0
	for i := 0; i < 1000; i++ {
		cur = cur.next
	}
	fmt.Println("1000", cur.val)
	res += cur.val
	for i := 0; i < 1000; i++ {
		cur = cur.next
	}
	fmt.Println("2000", cur.val)
	res += cur.val
	for i := 0; i < 1000; i++ {
		cur = cur.next
	}
	fmt.Println("3000", cur.val)
	res += cur.val
	return res
}
