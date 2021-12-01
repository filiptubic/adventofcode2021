package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadInput(path string) ([]int, error) {
	input := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, num)
	}
	return input, nil
}

func calculateWindow(input []int, index int) int {
	if len(input) <= index+2 {
		return -1
	}
	return input[index] + input[index+1] + input[index+2]
}

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	counter := 0
	last := calculateWindow(input, 0)
	for i := 1; i < len(input); i++ {
		window := calculateWindow(input, i)
		if window == -1 {
			break
		}

		if window > last {
			counter++
		}
		last = window
	}
	fmt.Println(counter)
}
