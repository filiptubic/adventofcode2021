package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value               int
	parent, left, right *Node
}

func (n Node) String() string {
	if n.left == nil && n.right == nil {
		return fmt.Sprintf("%d", n.value)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func makeTree(text string, parent *Node) *Node {
	if strings.HasPrefix(text, "[") && strings.HasSuffix(text, "]") {
		text = text[1 : len(text)-1]
		index := -1
		open := 0
		for i, r := range text {
			if r == ',' && open == 0 {
				index = i
				break
			}
			if r == '[' {
				open++
			} else if r == ']' {
				open--
			}
		}
		if index == -1 {
			panic("index == -1")
		}

		left, right := text[:index], text[index+1:]
		newNode := &Node{
			parent: parent,
			value:  math.MaxInt,
		}
		newNode.left = makeTree(left, newNode)
		newNode.right = makeTree(right, newNode)

		return newNode
	}
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return &Node{
		parent: parent,
		value:  num,
	}
}

func loadInput(path string) []*Node {
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	trees := make([]*Node, 0)

	for scanner.Scan() {
		line := scanner.Text()
		tree := makeTree(line, nil)
		trees = append(trees, tree)
	}
	return trees
}

func isLeaf(n *Node) bool {
	if n.left == nil && n.right == nil {
		return true
	}
	return false
}

func isFinalSubTree(n *Node) bool {
	if n == nil {
		return false
	}
	if n.left != nil && n.right != nil && isLeaf(n.left) && isLeaf(n.right) {
		return true
	}
	return false
}

func findNeighborLeft(n *Node, visited map[*Node]struct{}) *Node {
	if n == nil {
		return nil
	}

	for {
		if _, ok := visited[n.left]; ok {
			visited[n] = struct{}{}
			if n.parent == nil {
				return nil
			}
			n = n.parent
			continue
		}
		break
	}

	n = n.left
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func findNeighborRight(n *Node, visited map[*Node]struct{}) *Node {
	if n == nil {
		return nil
	}

	for {
		if _, ok := visited[n.right]; ok {
			visited[n] = struct{}{}
			if n.parent == nil {
				return nil
			}
			n = n.parent
			continue
		}
		break
	}
	n = n.right
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func explode(n *Node, depth int) bool {
	if isFinalSubTree(n) && depth >= 4 {
		visited := map[*Node]struct{}{n: {}}
		left := findNeighborLeft(n.parent, visited)
		if left != nil {
			left.value += n.left.value
		}
		visited = map[*Node]struct{}{n: {}}
		right := findNeighborRight(n.parent, visited)
		if right != nil {
			right.value += n.right.value
		}
		n.value = 0
		n.left = nil
		n.right = nil
		return true
	}

	exploded := false
	if n.left != nil {
		exploded = explode(n.left, depth+1)
	}
	if !exploded && n.right != nil {
		exploded = explode(n.right, depth+1)
	}
	return exploded
}

func split(n *Node, depth int) bool {
	if n == nil {
		return false
	}

	if isLeaf(n) && n.value >= 10 {
		newLeft := &Node{value: int(math.Floor(float64(n.value) / 2))}
		newRight := &Node{value: int(math.Ceil(float64(n.value) / 2))}
		newNode := &Node{
			left:  newLeft,
			right: newRight,
		}
		newLeft.parent = newNode
		newRight.parent = newNode

		if n.parent.left == n {
			n.parent.left = newNode
			newNode.parent = n.parent
		} else {
			n.parent.right = newNode
			newNode.parent = n.parent
		}
		return true
	}

	splited := split(n.left, depth+1)
	if !splited {
		splited = split(n.right, depth+1)
	}
	return splited
}

func copyGraph(n *Node, parent *Node) *Node {
	if n == nil {
		return nil
	}

	newN := &Node{
		parent: parent,
		value:  n.value,
	}
	newN.left = copyGraph(n.left, newN)
	newN.right = copyGraph(n.right, newN)

	return newN

}

func add(n1, n2 *Node) *Node {
	sum := &Node{
		left:  n1,
		right: n2,
	}
	sum.left.parent = sum
	sum.right.parent = sum

	for {
		exploded := explode(sum, 0)
		if exploded {
			continue
		}
		splitted := split(sum, 0)

		if !exploded && !splitted {
			break
		}
	}

	return sum
}

func magnitude(n *Node) int {
	if n == nil {
		return 0
	}
	if isFinalSubTree(n) {
		return 3*n.left.value + 2*n.right.value
	}
	return 3*magnitude(n.left) + 2*magnitude(n.right)
}

func main() {
	trees := loadInput("input.txt")
	treesCopy := make([]*Node, len(trees))
	for i := 0; i < len(trees); i++ {
		treesCopy[i] = copyGraph(trees[i], nil)
	}

	sum := treesCopy[0]
	for _, t := range treesCopy[1:] {
		sum = add(sum, t)
	}
	fmt.Printf("Part1: %d\n", magnitude(sum))

	max := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees); j++ {
			if i != j {
				n1 := copyGraph(trees[i], nil)
				n2 := copyGraph(trees[j], nil)
				sum := add(n1, n2)
				m := magnitude(sum)
				if max < m {
					max = m
				}
			}
		}
	}
	fmt.Printf("Part2: %d\n", max)
}
