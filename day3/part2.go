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

func calculateRating(binaries []binary, cmp func(zeros, ones []binary) []binary) binary {
	selected := binaries
	iteration := 0
	for len(selected) != 1 {
		selected = selectBinaries(selected, iteration, cmp)
		iteration++
	}
	return selected[0]
}

func selectBinaries(binaries []binary, index int, cmp func(zeros, ones []binary) []binary) []binary {
	zeros := make([]binary, 0)
	ones := make([]binary, 0)

	for i := 0; i < len(binaries); i++ {
		switch binaries[i][index] {
		case false:
			zeros = append(zeros, binaries[i])
		default:
			ones = append(ones, binaries[i])
		}
	}
	return cmp(zeros, ones)
}

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	oxygen := calculateRating(input, func(zeros, ones []binary) []binary {
		if len(ones) >= len(zeros) {
			return ones
		} else {
			return zeros
		}
	})

	oxygenNum, err := oxygen.num()
	if err != nil {
		panic(err)
	}

	co2 := calculateRating(input, func(zeros, ones []binary) []binary {
		if len(zeros) <= len(ones) {
			return zeros
		} else {
			return ones
		}
	})
	co2Num, err := co2.num()
	if err != nil {
		panic(err)
	}

	fmt.Println(oxygenNum * co2Num)

}
