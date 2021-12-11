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
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	return matrix, nil
}

func explode(i, j int, matrix [][]int, lights [][]bool) {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[i]) {
		// out of bound
		return
	}
	if matrix[i][j] >= 9 {
		matrix[i][j] = 0
		lights[i][j] = true
		explode(i+1, j, matrix, lights)   // bottom
		explode(i-1, j, matrix, lights)   // top
		explode(i, j+1, matrix, lights)   // right
		explode(i, j-1, matrix, lights)   // left
		explode(i-1, j-1, matrix, lights) // top, left
		explode(i-1, j+1, matrix, lights) // top, right
		explode(i+1, j-1, matrix, lights) // bottom, left
		explode(i+1, j+1, matrix, lights) // bottom, right
	} else if !lights[i][j] {
		matrix[i][j]++
	}
}

func initLights(m [][]int) [][]bool {
	lights := make([][]bool, 0)
	for _, row := range m {
		lights = append(lights, make([]bool, len(row)))
	}
	return lights
}

func main() {
	var lights [][]bool
	matrix, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	count := 0
	for step := 0; step < 100; step++ {
		lights = initLights(matrix)

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				explode(i, j, matrix, lights)
			}
		}

		for _, r := range lights {
			for _, l := range r {
				if l {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
