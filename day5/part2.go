package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isDiagonal(x1, x2, y1, y2 int) bool {
	return math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2))
}

func drawLine(matrix [][]int, x1, y1, x2, y2 int) {
	if isDiagonal(x1, x2, y1, y2) {
		// diagonal
		for x1 != x2 && y1 != y2 {
			matrix[y1][x1]++
			if y1 > y2 {
				y1--
			} else if y1 < y2 {
				y1++
			}
			if x1 > x2 {
				x1--
			} else if x1 < x2 {
				x1++
			}
		}
		matrix[y1][x1]++
	} else if x1 == x2 {
		// vertical
		for y1 != y2 {
			matrix[y1][x1]++
			if y1 > y2 {
				y1--
			} else {
				y1++
			}
		}
		matrix[y1][x1]++
	} else if y1 == y2 {
		// horizontal
		for x1 != x2 {
			matrix[y1][x1]++
			if x1 > x2 {
				x1--
			} else {
				x1++
			}
		}
		matrix[y1][x1]++
	}

}

func loadMatrix(path string) ([][]int, error) {
	// allocate 1000x1000 matrix
	matrix := make([][]int, 1000)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 1000)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "->")

		point1 := strings.Trim(parts[0], " ")
		point2 := strings.Trim(parts[1], " ")

		point1Parts := strings.Split(point1, ",")
		point2Parts := strings.Split(point2, ",")

		x1, err := strconv.Atoi(strings.Trim(point1Parts[0], " "))
		if err != nil {
			return nil, err
		}
		y1, err := strconv.Atoi(strings.Trim(point1Parts[1], " "))
		if err != nil {
			return nil, err
		}
		x2, err := strconv.Atoi(strings.Trim(point2Parts[0], " "))
		if err != nil {
			return nil, err
		}
		y2, err := strconv.Atoi(strings.Trim(point2Parts[1], " "))
		if err != nil {
			return nil, err
		}
		drawLine(matrix, x1, y1, x2, y2)
	}
	return matrix, nil
}

func countOverlaps(m [][]int) int {
	count := 0
	for _, r := range m {
		for _, x := range r {
			if x > 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	m, err := loadMatrix("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(countOverlaps(m))
}
