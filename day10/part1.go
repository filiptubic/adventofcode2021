package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func isOpen(r rune) bool {
	return r == '(' || r == '[' || r == '{' || r == '<'
}

func isClose(r rune) bool {
	return !isClose(r)
}

func checkClosing(open, close rune) bool {
	if open == '(' && close == ')' {
		return true
	}
	if open == '{' && close == '}' {
		return true
	}
	if open == '[' && close == ']' {
		return true
	}
	if open == '<' && close == '>' {
		return true
	}
	return false
}

func checkChunk(current rune, rest []rune) (incorrect rune, left []rune) {
	for len(rest) > 0 {
		if isOpen(rest[0]) {
			incorrect, left := checkChunk(rest[0], rest[1:])
			if incorrect == -1 {
				rest = left
				continue
			} else {
				return incorrect, left
			}
		}

		if !checkClosing(current, rest[0]) {
			return rest[0], rest[1:]
		} else {
			return -1, rest[1:]
		}
	}
	return -1, []rune{}
}

func errorCount(r rune) (int, error) {
	switch r {
	case ')':
		return 3, nil
	case ']':
		return 57, nil
	case '}':
		return 1197, nil
	case '>':
		return 25137, nil
	default:
		return -1, fmt.Errorf("invalid rune %s", string(r))
	}
}

func main() {
	lines, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	totalCount := make(map[rune]int)

	for _, line := range lines {
		runes := []rune(line)
		illegal, _ := checkChunk(runes[0], runes[1:])
		if illegal != -1 {
			totalCount[illegal]++
		}
	}
	total := 0
	for key, count := range totalCount {
		ec, err := errorCount(key)
		if err != nil {
			panic(err)
		}
		total += ec * count
	}
	fmt.Println(total)
}
