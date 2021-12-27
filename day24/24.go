package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	index, value int
}

type Stack []Pair

func (s Stack) push(p Pair) Stack {
	return append(s, p)
}

func (s Stack) pop() (Pair, Stack) {
	last := s[len(s)-1]
	return last, s[:len(s)-1]
}

func solve() {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	stack := make(Stack, 0)
	digitsMin := make(map[int]int)
	digitsMax := make(map[int]int)
	lineNumber := 0

	var last Pair
	var push bool
	var sub int // used by pop part
	var currentDigit = 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")[1:]

		if lineNumber%18 == 4 {
			// div z 1 means its push part, otherwise is pop
			push = parts[1] == "1"
		}

		if lineNumber%18 == 5 {
			// keep this value in case pop part is active
			sub, err = strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
		}

		if lineNumber%18 == 15 {
			if push {
				value, err := strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}
				stack = stack.push(Pair{currentDigit, value})
			} else {
				last, stack = stack.pop()
				diff := last.value + sub
				if diff < 0 {
					digitsMin[last.index] = -diff + 1
					digitsMin[currentDigit] = 1

					digitsMax[last.index] = 9
					digitsMax[currentDigit] = 9 + diff
				} else {
					digitsMin[last.index] = 1
					digitsMin[currentDigit] = 1 + diff

					digitsMax[last.index] = 9 - diff
					digitsMax[currentDigit] = 9
				}
			}
			currentDigit++
		}
		lineNumber++
	}

	num := ""
	for i := 0; i < 14; i++ {
		num = fmt.Sprintf("%s%d", num, digitsMax[i])
	}
	fmt.Printf("Part 1: %s\n", num)

	num = ""
	for i := 0; i < 14; i++ {
		num = fmt.Sprintf("%s%d", num, digitsMin[i])
	}
	fmt.Printf("Part 2: %s\n", num)
}

func main() {
	solve()
}
