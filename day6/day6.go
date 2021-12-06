package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	numbers := make([]int, 0)
	for _, part := range strings.Split(line, ",") {
		if part == "" {
			continue
		}
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func evolve(number, current, depth int, done map[int]map[int]int) int {
	if m, ok := done[current]; ok {
		if value, ok := m[number]; ok {
			return value
		}
	} else {
		done[current] = make(map[int]int)
	}

	if current == depth {
		return 1
	}

	if number == 0 {
		done[current][number] = evolve(6, current+1, depth, done) + evolve(8, current+1, depth, done)
		return done[current][number]
	}

	done[current][number] = evolve(number-1, current+1, depth, done)
	return done[current][number]
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("missing argument days")
	}
	days, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	nums, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	populationCount := 0
	done := make(map[int]map[int]int)
	for _, num := range nums {
		populationCount += evolve(num, 0, days, done)
	}

	fmt.Println(populationCount)
}
