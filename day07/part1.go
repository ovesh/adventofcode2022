package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type dir struct {
	parent   *dir
	children []*dir
	files    []*file
	size     int
	name     string
}

func (n *dir) String() string {
	return fmt.Sprintf("%s (%d), files:\n[%v],\nchildren:\n[%v]", n.name, n.size, n.files, n.children)
}

type file struct {
	name string
	size int
}

func (f *file) String() string {
	return fmt.Sprintf("%s (%d)\n", f.name, f.size)
}

func newNode(parent *dir, name string, size int) *dir {
	return &dir{
		parent:   parent,
		children: make([]*dir, 0),
		files:    make([]*file, 0),
		size:     size,
		name:     name,
	}
}

func (n *dir) addChild(child *dir) {
	n.children = append(n.children, child)
	//n.size += child.size
}

func (n *dir) addFile(f *file) {
	n.size += f.size
	n.files = append(n.files, f)
	for curr := n.parent; curr != nil; curr = curr.parent {
		curr.size += f.size
	}
}

func (n *dir) findChild(name string) *dir {
	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}
	panic("not found")
}

func atoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func parseInput(input string) (*dir, []*dir) {
	lines := strings.Split(input, "\n")
	root := newNode(nil, "", 0)
	allNodes := []*dir{root}
	current := root
	inLS := false
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "$ ls") {
			inLS = true
			continue
		} else if strings.HasPrefix(line, "$ cd ") {
			inLS = false
			if strings.HasSuffix(line, "/") {
				current = root
			} else if strings.HasSuffix(line, "..") {
				// assume never `cd ..` from root
				current = current.parent
			} else {
				dirs := strings.Split(line[5:], "/")
				for _, d := range dirs {
					current = current.findChild(d)
				}
			}
			continue
		}

		if !inLS {
			continue
		}

		if strings.HasPrefix(line, "dir ") {
			// assume ls is never called twice on a dir
			node := newNode(current, line[4:], 0)
			allNodes = append(allNodes, node)
			current.addChild(node)
		} else {
			both := strings.Split(line, " ")
			size := both[0]
			name := both[1]
			current.addFile(&file{name: name, size: atoi(size)})
		}
	}

	return root, allNodes
}

func Part1() int {
	_, allNodes := parseInput(input0)

	total := 0
	for _, n := range allNodes {
		if n.size <= 100000 {
			total += n.size
		}
	}

	return total
}
