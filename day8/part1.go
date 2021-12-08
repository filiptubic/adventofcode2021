package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadInput(path string) ([][]string, [][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	allSignals := make([][]string, 0)
	allOutput := make([][]string, 0)
	for scanner.Scan() {
		signals := make([]string, 0)
		output := make([]string, 0)

		lineParts := strings.Split(scanner.Text(), "|")
		signalsPart := strings.Trim(lineParts[0], " ")
		outputPart := strings.Trim(lineParts[1], " ")

		for _, signal := range strings.Split(signalsPart, " ") {
			signal = strings.Trim(signal, " ")
			if signal != "" {
				signals = append(signals, signal)
			}
		}
		for _, out := range strings.Split(outputPart, " ") {
			out = strings.Trim(out, " ")
			if out != "" {
				output = append(output, out)
			}
		}
		allSignals = append(allSignals, signals)
		allOutput = append(allOutput, output)
	}

	return allSignals, allOutput, nil
}

func countUnique(output []string) int {
	c := 0
	for _, o := range output {
		n := len(o)
		if n == 2 || n == 4 || n == 3 || n == 7 {
			c++
		}
	}
	return c
}

func main() {
	_, allOutput, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	c := 0
	for _, output := range allOutput {
		c += countUnique(output)
	}
	fmt.Println(c)
}
