package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	i, j int
}

func loadCucumbers(path string) (cucumbers [][]rune, nrows, ncols int) {
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(fp)
	s.Split(bufio.ScanLines)
	nrows = 0
	ncols = 0
	cucumbers = make([][]rune, 0)
	for s.Scan() {
		line := s.Text()
		ncols = len(line)
		cucumbers = append(cucumbers, []rune(line))
		nrows++
	}
	return cucumbers, nrows, ncols
}

func move(cucumbers [][]rune, nrows, ncols int, dir rune, offset Position) bool {
	moves := make(map[Position]Position)

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			if cucumbers[i][j] == dir {
				destination := Position{i: (i + offset.i) % nrows, j: (j + offset.j) % ncols}
				if cucumbers[destination.i][destination.j] == '.' {
					moves[Position{i, j}] = destination
				}
			}
		}
	}

	if len(moves) > 0 {
		for src, dst := range moves {
			cucumbers[src.i][src.j] = '.'
			cucumbers[dst.i][dst.j] = dir
		}
		return true
	}

	return false
}

func solve(cucumbers [][]rune, nrows, ncols int) (steps int) {
	right := true
	down := true
	for right || down {
		steps++
		right = move(cucumbers, nrows, ncols, '>', Position{0, 1})
		down = move(cucumbers, nrows, ncols, 'v', Position{1, 0})
	}
	return steps
}

func main() {
	cucumbers, nrows, ncols := loadCucumbers("input.txt")
	fmt.Println(solve(cucumbers, nrows, ncols))
}
