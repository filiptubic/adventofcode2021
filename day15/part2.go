package main

import (
	"bufio"
	"container/heap"
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

	// extend right
	mOriginalSize := len(m[0])

	for i := 0; i < len(m); i++ {
		for size := 0; size < 4; size++ {
			for j := size * mOriginalSize; j < size*mOriginalSize+mOriginalSize; j++ {
				x := m[i][j] + 1
				if x > 9 {
					x = 1
				}
				m[i] = append(m[i], x)
			}
		}
	}

	// extend bottom
	mOriginalSize = len(m)
	for size := 0; size < 4; size++ {
		for i := 0; i < mOriginalSize; i++ {
			m = append(m, make([]int, len(m[0])))
		}
	}

	for i := mOriginalSize; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			x := m[i-mOriginalSize][j] + 1
			if x > 9 {
				x = 1
			}
			m[i][j] = x
		}
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

// An Item is something we manage in a priority queue.
type Item struct {
	value    *Node
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value *Node, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func dijkstraPriorityQueue(nodes []*Node, weights map[*Node]map[*Node]int) (map[*Node]int, map[*Node]*Node) {
	dist := make(map[*Node]int)
	prev := make(map[*Node]*Node)
	dist[nodes[0]] = 0

	pq := make(PriorityQueue, len(nodes))

	pq[0] = &Item{
		value:    nodes[0],
		priority: 0,
	}
	pqMap := make(map[*Node]*Item)
	pqMap[nodes[0]] = pq[0]

	for idx, node := range nodes[1:] {
		dist[node] = math.MaxInt64
		pq[idx+1] = &Item{
			value:    node,
			priority: math.MaxInt64,
		}
		pqMap[node] = pq[idx+1]
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.value

		for v, cost := range weights[u] {
			alt := math.MaxInt64
			if dist[u] != math.MaxInt64 {
				alt = dist[u] + cost
			}

			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				heap.Push(&pq, &Item{value: v, priority: dist[v]})
			}
		}
	}

	return dist, prev
}

func main() {
	m, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	nodes, weights := makeGraph(m)
	last := nodes[len(nodes)-1]
	dist, _ := dijkstraPriorityQueue(nodes, weights)

	fmt.Println(dist[last])
}
