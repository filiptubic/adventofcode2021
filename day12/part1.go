package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func loadInput(path string) (map[string][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	m := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		left, right := parts[0], parts[1]

		if val, ok := m[left]; ok {
			m[left] = append(val, right)
		} else {
			m[left] = []string{right}
		}
		if val, ok := m[right]; ok {
			m[right] = append(val, left)
		} else {
			m[right] = []string{left}
		}
	}
	return m, nil
}

func isUpper(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

func isVisited(visited []string, node string) bool {
	if isUpper(node) {
		return false
	}

	for _, v := range visited {
		if v == node {
			return true
		}
	}
	return false
}

func findPaths(node string, m map[string][]string, visited []string) []string {
	if node == "end" {
		return []string{"end"}
	}

	paths := make([]string, 0)
	for _, n := range m[node] {
		if !isVisited(visited, n) {
			for _, newPath := range findPaths(n, m, append(visited, node)) {
				paths = append(paths, fmt.Sprintf("%s%s", node, newPath))
			}
		}
	}
	return paths
}

func main() {
	m, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	paths := findPaths("start", m, make([]string, 0))
	fmt.Println(len(paths))
}
