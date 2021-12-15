package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	return !isOpen(r)
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

func scoreOf(r rune) (int, error) {
	switch r {
	case ')':
		return 1, nil
	case ']':
		return 2, nil
	case '}':
		return 3, nil
	case '>':
		return 4, nil
	default:
		return -1, fmt.Errorf("unsupported rune %s", string(r))
	}
}

func invert(r rune) (rune, error) {
	switch r {
	case ')':
		return '(', nil
	case '(':
		return ')', nil
	case ']':
		return '[', nil
	case '[':
		return ']', nil
	case '}':
		return '{', nil
	case '{':
		return '}', nil
	case '>':
		return '<', nil
	case '<':
		return '>', nil
	default:
		return -1, fmt.Errorf("invalid rune %s to invert", string(r))
	}
}

func fixChunk(chunk []rune) (map[rune]int, []rune, error) {
	if len(chunk) == 0 {
		return make(map[rune]int), []rune{}, nil
	}

	m, missing, err := fixChunk(chunk[1:])
	if err != nil {
		return m, missing, err
	}

	if isClose(chunk[0]) {
		m[chunk[0]]++
		return m, missing, nil
	}

	inverted, err := invert(chunk[0])
	if err != nil {
		return m, missing, err
	}

	if m[inverted] > 0 {
		m[inverted]--
		return m, missing, nil
	}
	missing = append(missing, inverted)
	return m, missing, nil
}

func main() {
	lines, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	incomplete := make([]string, 0)
	for _, line := range lines {
		runes := []rune(line)
		illegal, _ := checkChunk(runes[0], runes[1:])
		if illegal == -1 {
			incomplete = append(incomplete, line)
		}
	}

	totalScores := make([]int, 0)
	for _, l := range incomplete {
		_, missing, err := fixChunk([]rune(l))
		if err != nil {
			panic(err)
		}

		total := 0
		for _, r := range missing {
			s, err := scoreOf(r)
			if err != nil {
				panic(err)
			}
			total = total*5 + s
		}
		totalScores = append(totalScores, total)
	}
	sort.Ints(totalScores)
	fmt.Println(totalScores[len(totalScores)/2])
}
