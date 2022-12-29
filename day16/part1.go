package day16

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	structs "github.com/kirves/godijkstra/common/structs"
	"github.com/kirves/godijkstra/dijkstra"
)

func atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

type node struct {
	id        string
	rotated   bool
	flowRate  int
	neighbors []*node
}

type GraphObject struct {
	nodes2Neighbors map[string][]structs.Connection
}

func (g *GraphObject) SuccessorsForNode(node string) []structs.Connection {
	return g.nodes2Neighbors[node]
}

func (g *GraphObject) PredecessorsFromNode(_ string) []structs.Connection {
	return []structs.Connection{}
}
func (g *GraphObject) EdgeWeight(_, _ string) int {
	return 1
}

// all shorted paths between nodes that have a flow rate > 0 (and from the first node)
func findAllShortestRoutes(nodes []*node, aa *node) map[string]int {
	graph := &GraphObject{nodes2Neighbors: map[string][]structs.Connection{}}
	for _, node := range nodes {
		nodeID := node.id
		graph.nodes2Neighbors[nodeID] = []structs.Connection{}
		for _, neighbor := range node.neighbors {
			graph.nodes2Neighbors[nodeID] = append(graph.nodes2Neighbors[nodeID], structs.Connection{Destination: neighbor.id, Weight: 1})
		}
	}
	res := map[string]int{}
	for _, nodeA := range nodes {
		for _, nodeB := range nodes {
			if (nodeA.flowRate == 0 || nodeB.flowRate == 0) && (nodeA != aa && nodeB != aa) {
				continue
			}
			if nodeA == nodeB {
				continue
			}
			path, valid := dijkstra.SearchPath(graph, nodeA.id, nodeB.id, dijkstra.VANILLA)
			if !valid {
				panic("invalid path")
			}
			//fmt.Println(path.Path)
			res[nodeA.id+","+nodeB.id] = len(path.Path) - 1
		}
	}
	return res
}

func (n *node) String() string {
	res := fmt.Sprintf("%v (%v): ", n.id, n.flowRate)
	for _, neighbor := range n.neighbors {
		res += fmt.Sprintf("%v, ", neighbor.id)
	}
	return res
}

type state struct {
	closed        map[string]struct{}
	current       string
	remainingTime int
	score         int
}

func (s state) String() string {
	closed := make([]string, len(s.closed))
	i := 0
	for k := range s.closed {
		closed[i] = k
		i++
	}
	sort.Strings(closed)
	return fmt.Sprintf("%v %v %v %v", closed, s.current, s.remainingTime, s.score)
}

var seenStates = map[string]struct{}{}

func next(states []state, shortestRoutes map[string]int, nodesByID map[string]*node) []state {
	res := []state{}
	for _, st := range states {
		if _, ok := seenStates[st.String()]; ok {
			continue
		}
		seenStates[st.String()] = struct{}{}
		if st.remainingTime == 0 {
			continue
		}
		if len(st.closed) == 0 {
			nextState := state{closed: st.closed, current: st.current, remainingTime: 0, score: st.score}
			res = append(res, nextState)
		}
		if _, ok := st.closed[st.current]; ok {
			newClosed := map[string]struct{}{}
			for id := range st.closed {
				if id == st.current {
					continue
				}
				newClosed[id] = struct{}{}
			}
			nextState := state{closed: newClosed, current: st.current, remainingTime: st.remainingTime - 1, score: st.score + (st.remainingTime-1)*nodesByID[st.current].flowRate}
			res = append(res, nextState)
		}
		for nodeID := range st.closed {
			if st.current == nodeID {
				continue
			}
			shortest, ok := shortestRoutes[st.current+","+nodeID]
			if !ok {
				panic(fmt.Sprintf("no shortest route for %q\n", st.current+","+nodeID))
			}
			newRemainingTime := st.remainingTime - shortest
			if newRemainingTime >= 0 {
				nextState := state{closed: st.closed, current: nodeID, remainingTime: newRemainingTime, score: st.score}
				res = append(res, nextState)
			}
		}
	}
	return res
}

func Part1() int {
	lines := strings.Split(input0, "\n")
	nodeByID := make(map[string]*node)
	nodes := make([]*node, len(lines))
	nodesWithFlow := map[string]struct{}{}
	var aa *node
	for i, line := range lines {
		valveID := line[len("Valve ") : len("Valve ")+2]
		flowRate := atoi(line[len("Valve AA has flow rate="):strings.Index(line, ";")])
		if flowRate > 0 {
			nodesWithFlow[valveID] = struct{}{}
		}
		nodeByID[valveID] = &node{id: valveID, flowRate: flowRate, neighbors: nil}
		nodes[i] = nodeByID[valveID]
		if valveID == "AA" {
			aa = nodes[i]
		}
	}
	fmt.Println(nodeByID)
	for i, line := range lines {
		to := strings.Split(line[strings.Index(line, ";")+len("; tunnels lead to valves "):], ", ")
		for _, neighborID := range to {
			neighbor, ok := nodeByID[neighborID]
			if !ok {
				panic(fmt.Sprintf("%q", neighborID))
			}
			nodes[i].neighbors = append(nodes[i].neighbors, neighbor)
		}
	}

	return nodes, nodesWithFlow, nodeByID, aa
}

func Part1() int {
	nodes, nodesWithFlow, nodeByID, aa := parseInput(input0)
	shortestRoutes := findAllShortestRoutes(nodes, aa)
	fmt.Println(shortestRoutes)
	max := 0
	for states := []state{{closed: nodesWithFlow, current: aa.id, remainingTime: 30, score: 0}}; len(states) > 0; states = next(states, shortestRoutes, nodeByID) {
		fmt.Println("len(states)", len(states))
		for _, st := range states {
			//fmt.Println("st.remainingTime, st.score", st.remainingTime, st.score)
			if st.remainingTime == 0 {
				if st.score > max {
					fmt.Println("new max", st.score)
					max = st.score
				}
			}
		}
	}
	return max
}
