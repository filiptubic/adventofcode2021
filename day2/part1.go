package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type action struct {
	path  string
	steps int
}

func loadInput(path string) ([]action, error) {
	input := make([]action, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		path, steps := parts[0], parts[1]

		stepsNum, err := strconv.Atoi(steps)
		if err != nil {
			return nil, err
		}
		input = append(input, action{path: path, steps: stepsNum})
	}
	return input, nil
}

func main() {
	actions, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	x, depth := 0, 0
	for _, action := range actions {
		switch action.path {
		case "forward":
			x += action.steps
		case "down":
			depth += action.steps
		case "up":
			depth -= action.steps
		}
	}
	fmt.Println(x * depth)
}
