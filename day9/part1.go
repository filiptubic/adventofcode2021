package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	m, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	risk := 0
	for i := 0; i < len(m); i++ {
		lowPoints := make([]int, 0)
		for j := 0; j < len(m[i]); j++ {
			if lowPoint(i, j, m) {
				lowPoints = append(lowPoints, m[i][j])
				risk = risk + 1 + m[i][j]
			}
		}
	}
	fmt.Println(risk)
}
