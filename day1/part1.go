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

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	counter := 0
	last := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > last {
			counter++
		}
		last = input[i]
	}
	fmt.Println(counter)
}
