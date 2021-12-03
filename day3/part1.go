package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type binary []bool

func (b binary) String() string {
	output := ""
	for _, v := range b {
		switch v {
		case true:
			output = fmt.Sprintf("%s1", output)
		default:
			output = fmt.Sprintf("%s0", output)
		}
	}
	return output
}

func (b binary) invert() binary {
	inverted := binary{}
	for _, v := range b {
		inverted = append(inverted, !v)
	}
	return inverted
}

func (b binary) num() (int64, error) {
	parsedNum, err := strconv.ParseInt(b.String(), 2, 64)
	if err != nil {
		return -1, err
	}
	return parsedNum, nil
}

func calulateGamma(binaries []binary) binary {
	gamma := binary{}

	for i := 0; i < len(binaries[0]); i++ {
		counterOne, counterZero := 0, 0
		for j := 0; j < len(binaries); j++ {
			switch binaries[j][i] {
			case true:
				counterOne++
			default:
				counterZero++
			}
		}
		if counterOne > counterZero {
			gamma = append(gamma, true)
		} else {
			gamma = append(gamma, false)
		}
	}
	return gamma
}

func loadInput(path string) ([]binary, error) {
	input := make([]binary, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		b := make(binary, 0)
		for _, r := range line {
			switch r {
			case '0':
				b = append(b, false)
			default:
				b = append(b, true)
			}
		}
		input = append(input, b)
	}
	return input, nil
}

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	gamma := calulateGamma(input)
	gammaNum, err := gamma.num()
	if err != nil {
		panic(err)
	}
	epsilonNum, err := gamma.invert().num()
	if err != nil {
		panic(err)
	}
	fmt.Println(gammaNum * epsilonNum)
}
