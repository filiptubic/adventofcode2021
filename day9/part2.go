package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func loadInput(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	matrix := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		line := scanner.Text()
		for _, r := range line {
			number, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			row = append(row, number)
		}

		matrix = append(matrix, row)
	}
	return matrix, nil
}

func lowPoint(i int, j int, m [][]int) bool {
	if j-1 >= 0 {
		// left
		if m[i][j] >= m[i][j-1] {
			return false
		}
	}
	if j+1 < len(m[i]) {
		// right
		if m[i][j] >= m[i][j+1] {
			return false
		}
	}
	if i-1 >= 0 {
		// top
		if m[i][j] >= m[i-1][j] {
			return false
		}
	}
	if i+1 < len(m) {
		// bottom
		if m[i][j] >= m[i+1][j] {
			return false
		}
	}
	return true
}

func basin(i, j int, m [][]int, visited map[int]map[int]bool) int {
	if i < 0 || j < 0 || i >= len(m) || j >= len(m[i]) {
		// out of bound
		return 0
	}
	if m[i][j] == 9 {
		// skip 9
		return 0
	}

	if val, ok := visited[i][j]; val && ok {
		// already visited position (i,j)
		return 0
	}

	visited[i][j] = true

	top := basin(i-1, j, m, visited)
	bottom := basin(i+1, j, m, visited)
	left := basin(i, j-1, m, visited)
	right := basin(i, j+1, m, visited)
	return 1 + top + bottom + left + right
}

func initVisited(m [][]int) map[int]map[int]bool {
	visited := make(map[int]map[int]bool)
	for i := 0; i < len(m); i++ {
		visited[i] = make(map[int]bool)
		for j := 0; j < len(m[i]); j++ {
			visited[i][j] = false
		}
	}
	return visited
}

func main() {
	m, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	basins := make([]int, 0)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if lowPoint(i, j, m) {
				basins = append(basins, basin(i, j, m, initVisited(m)))
			}
		}
	}
	sort.Ints(basins)
	result := 1
	for _, b := range basins[len(basins)-3:] {
		result *= b
	}
	fmt.Println(result)
}
