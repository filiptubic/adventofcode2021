package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func loadInput(path string) (string, map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	template := scanner.Text()
	m := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		left, right := strings.Trim(parts[0], " "), strings.Trim(parts[1], " ")
		m[left] = right
	}
	return template, m, nil
}

func copyMap(m map[string]int) map[string]int {
	newM := make(map[string]int)
	for k, v := range m {
		newM[k] = v
	}
	return newM
}

func findCommon(template string, rules map[string]string, steps int) (int, int) {
	countMap := make(map[string]int)
	letterCount := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		left := template[i]
		right := template[i+1]
		letterCount[string(left)]++

		pair := string([]byte{left, right})
		if _, ok := rules[pair]; ok {
			countMap[pair]++
		}
	}
	letterCount[string(template[len(template)-1])]++

	for i := 0; i < steps; i++ {
		newCountMap := copyMap(countMap)

		for pair := range countMap {
			if countMap[pair] == 0 {
				continue
			}

			if newLetter, ok := rules[pair]; ok {
				pair1 := string(pair[0]) + newLetter
				pair2 := newLetter + string(pair[1])
				newCountMap[pair1] += countMap[pair]
				newCountMap[pair2] += countMap[pair]
				letterCount[newLetter] += countMap[pair]
			}
			newCountMap[pair] -= countMap[pair]

		}
		countMap = newCountMap
	}

	return findMaxMin(letterCount)
}

func findMaxMin(m map[string]int) (int, int) {
	min, max := math.MaxInt, math.MinInt
	for _, count := range m {
		if min > count {
			min = count
		}
		if max < count {
			max = count
		}
	}
	return max, min
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("missing steps argument")
		os.Exit(1)
	}

	steps, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	template, rules, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	max, min := findCommon(template, rules, steps)
	fmt.Println(max - min)
}
