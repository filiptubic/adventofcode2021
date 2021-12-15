package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FoldType int

const (
	X FoldType = iota
	Y
)

type Fold struct {
	Type  FoldType
	Value int
}

func maxNumber(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}

	}
	return max
}

func printMatrix(m [][]int, startX, startY, endX, endY int) {
	for i := startY; i <= endY; i++ {
		for j := startX; j <= endX; j++ {
			if m[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Println()
	}
}

func loadInput(path string) ([][]int, []Fold, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	allX := make([]int, 0)
	allY := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ",")
		left, right := parts[0], parts[1]
		x, err := strconv.Atoi(left)
		if err != nil {
			return nil, nil, err
		}
		y, err := strconv.Atoi(right)
		if err != nil {
			return nil, nil, err
		}

		allX = append(allX, x)
		allY = append(allY, y)
	}

	nrows := maxNumber(allY) + 1
	ncols := maxNumber(allX) + 1

	m := make([][]int, nrows)
	for i := 0; i < nrows; i++ {
		m[i] = make([]int, ncols)
	}

	for i := 0; i < len(allX); i++ {
		m[allY[i]][allX[i]] = 1
	}

	folds := make([]Fold, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if strings.HasPrefix(line, "fold along y=") {
			y := strings.TrimPrefix(line, "fold along y=")
			foldY, err := strconv.Atoi(y)
			if err != nil {
				return nil, nil, err
			}
			folds = append(folds, Fold{Type: Y, Value: foldY})
		}
		if strings.HasPrefix(line, "fold along x=") {
			x := strings.TrimPrefix(line, "fold along x=")
			foldX, err := strconv.Atoi(x)
			if err != nil {
				return nil, nil, err
			}
			folds = append(folds, Fold{Type: X, Value: foldX})
		}
	}

	return m, folds, nil
}

func rotateHorizontal(m [][]int, startY, endY int) {
	sizeM := endY - startY + 1
	for i := 0; i < sizeM/2; i++ {
		tmp := m[startY+i]
		m[startY+i] = m[startY+sizeM-i-1]
		m[startY+sizeM-i-1] = tmp
	}
}

func rotateVertical(m [][]int, startX, startY, endX, endY int) {
	sizeX := endX - startX + 1
	sizeY := endY - startY + 1
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX/2; j++ {
			tmp := m[startY+i][startX+j]
			m[startY+i][startX+j] = m[startY+i][endX-j]
			m[startY+i][endX-j] = tmp
		}
	}
}

func foldMatrix(m [][]int, startX, startY, endX, endY int, f Fold) (newStartX, newStartY, newEndX, newEndY int) {
	if f.Type == X {
		sizeY := (endY - startY) + 1
		sizeX := (endX - startX) + 1

		left := f.Value
		right := sizeX - left - 1

		if left > right {
			// fold to left
			for j := startY; j < startY+sizeY; j++ {
				for i := 0; i < right; i++ {
					if m[j][startX+f.Value+i+1] == 1 {
						m[j][startX+f.Value-i-1] = 1
					}
				}
			}
			endX -= (right + 1)
		} else {
			// fold to right
			for j := startY; j < startY+sizeY; j++ {
				for i := 0; i < left; i++ {
					if m[j][startX+f.Value-i-1] == 1 {
						m[j][startX+f.Value+i+1] = 1
					}
				}
			}
			startX += (left + 1)
			// rotate matrix vertically in order to be folded to left
			rotateVertical(m, startX, startY, endX, endY)
		}
	} else if f.Type == Y {
		sizeY := (endY - startY) + 1
		sizeX := (endX - startX) + 1

		top := f.Value
		bottom := sizeY - top - 1

		if top > bottom {
			// fold to top
			for j := startX; j < startX+sizeX; j++ {
				for i := 0; i < bottom; i++ {
					if m[startY+f.Value+i+1][j] == 1 {
						m[startY+f.Value-i-1][j] = 1
					}
				}
			}
			endY -= (bottom + 1)
		} else {
			// fold to bottom
			for j := startX; j < startX+sizeX; j++ {
				for i := 0; i < top; i++ {
					if m[startY+f.Value-i-1][j] == 1 {
						m[startY+f.Value+i+1][j] = 1
					}
				}
			}
			startY += (top + 1)
			// rotate matrix horizontally in order to be folded to top
			rotateHorizontal(m, startY, endY)
		}
	}
	return startX, startY, endX, endY
}

func main() {
	m, folds, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	startX := 0
	startY := 0
	endX := len(m[0]) - 1
	endY := len(m) - 1

	for _, f := range folds {
		startX, startY, endX, endY = foldMatrix(m, startX, startY, endX, endY, f)
	}

	counter := 0
	for i := startY; i < endY+1; i++ {
		for j := startX; j < endX+1; j++ {
			if m[i][j] == 1 {
				counter++
			}
		}
	}
	printMatrix(m, startX, startY, endX, endY)
	fmt.Println(counter)
}
