package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	i, j int
}

func (n *Node) String() string {
	return fmt.Sprintf("(%d, %d)", n.i, n.j)
}

func loadInput(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	m := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		nums := make([]int, 0)
		for _, r := range line {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		m = append(m, nums)
	}

	return m, nil
}

func makeGraph(m [][]int) ([]*Node, map[*Node]map[*Node]int) {
	nodes := make([]*Node, 0)

	nodeM := make([][]*Node, 0)
	for i := 0; i < len(m); i++ {
		row := make([]*Node, 0)
		for j := 0; j < len(m[i]); j++ {
			nodes = append(nodes, &Node{i: i, j: j})
			row = append(row, nodes[len(nodes)-1])
		}
		nodeM = append(nodeM, row)
	}

	weights := make(map[*Node]map[*Node]int)

	for i := 0; i < len(nodes); i++ {
		if weights[nodes[i]] == nil {
			weights[nodes[i]] = make(map[*Node]int)
		}

		if nodes[i].j+1 < len(m[nodes[i].i]) {
			weightRight := m[nodes[i].i][nodes[i].j+1]
			to := nodeM[nodes[i].i][nodes[i].j+1]
			weights[nodes[i]][to] = weightRight
		}
		if nodes[i].j-1 >= 0 {
			weightLeft := m[nodes[i].i][nodes[i].j-1]
			to := nodeM[nodes[i].i][nodes[i].j-1]
			weights[nodes[i]][to] = weightLeft
		}

		if nodes[i].i-1 >= 0 {
			weightTop := m[nodes[i].i-1][nodes[i].j]
			to := nodeM[nodes[i].i-1][nodes[i].j]
			weights[nodes[i]][to] = weightTop
		}
		if nodes[i].i+1 < len(m) {
			weightBottom := m[nodes[i].i+1][nodes[i].j]
			to := nodeM[nodes[i].i+1][nodes[i].j]
			weights[nodes[i]][to] = weightBottom
		}
	}
	return nodes, weights
}

func isIn(node *Node, nodes []*Node) bool {
	for i := 0; i < len(nodes); i++ {
		if node == nodes[i] {
			return true
		}
	}
	return false
}

func dijkstraList(nodes []*Node, weights map[*Node]map[*Node]int) map[*Node]int {
	q := make([]*Node, 0)
	dist := make(map[*Node]int)
	prev := make(map[*Node]*Node)

	for i := 0; i < len(nodes); i++ {
		dist[nodes[i]] = math.MaxInt64
		prev[nodes[i]] = nil
		q = append(q, nodes[i])
	}
	dist[nodes[0]] = 0

	for len(q) > 0 {
		u := q[0]
		minIndex := 0
		for i := 1; i < len(q); i++ {
			if dist[q[i]] < dist[u] {
				u = q[i]
				minIndex = i
			}
		}
		if len(q) > 1 && (minIndex+1 < len(q)) {
			q = append(q[:minIndex], q[minIndex+1:]...)
		} else {
			q = q[:minIndex]
		}

		for v, cost := range weights[u] {
			if !isIn(v, q) {
				continue
			}

			alt := dist[u] + cost

			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	return dist
}

func main() {
	m, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	nodes, weights := makeGraph(m)
	last := nodes[len(nodes)-1]
	dist := dijkstraList(nodes, weights)

	fmt.Println(dist[last])
}
