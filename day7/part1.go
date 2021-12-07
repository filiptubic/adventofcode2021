package main

import (
	"bufio"
	"fmt"
	"math"
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

func findMax(array []int) int {
	m := array[0]

	for i := 1; i < len(array); i++ {
		if m < array[i] {
			m = array[i]
		}
	}
	return m
}

func findMin(array []int) int {
	m := array[0]

	for i := 1; i < len(array); i++ {
		if m > array[i] {
			m = array[i]
		}
	}
	return m
}

func fuel(nums []int) int {
	max := findMax(nums)
	fuel := make([]int, max+1)

	for _, num := range nums {
		// right of num
		for i := num + 1; i < len(fuel); i++ {
			for j := 1; j <= int(math.Abs(float64(num-i))); j++ {
				fuel[i] += 1
			}
		}
		// left of num
		for i := num - 1; i >= 0; i-- {
			for j := 1; j <= int(math.Abs(float64(num-i))); j++ {
				fuel[i] += 1
			}
		}
	}

	return findMin(fuel)
}

func main() {
	nums, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(fuel(nums))
}
