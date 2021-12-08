package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func contains(s string, r rune) bool {
	for _, elem := range s {
		if elem == r {
			return true
		}
	}
	return false
}

func containedIn(s1, s2 string) bool {
	for _, r := range s1 {
		if !contains(s2, r) {
			return false
		}
	}
	return true
}

func concat(s1, s2 string) string {
	for _, r := range s2 {
		if !contains(s1, r) {
			s1 = fmt.Sprintf("%s%s", s1, string(r))
		}
	}
	return s1
}

func makeDigits(digits []string) []string {
	nums := make([]string, 10)
	rest := make([]string, 0)

	for _, d := range digits {
		switch len(d) {
		case 2:
			nums[1] = d
		case 4:
			nums[4] = d
		case 3:
			nums[7] = d
		case 7:
			nums[8] = d
		default:
			rest = append(rest, d)
		}
	}
	digits = rest
	rest = make([]string, 0)

	// 6
	tmp := nums[8]
	for _, r := range nums[1] {
		if contains(tmp, r) {
			tmp = strings.Replace(tmp, string(r), "", -1)
		}
	}

	found6 := false
	for _, d := range digits {
		if !found6 && len(d) == 6 && containedIn(tmp, d) {
			nums[6] = d
			found6 = true
			continue
		}
		rest = append(rest, d)
	}
	digits = rest
	rest = make([]string, 0)

	// 9
	found9 := false
	tmp = concat(nums[4], nums[7])
	for _, d := range digits {
		if !found9 && len(d) == 6 && containedIn(tmp, d) {
			nums[9] = d
			found9 = true
			continue
		}
		rest = append(rest, d)
	}
	digits = rest
	rest = make([]string, 0)

	// 0
	found0 := false
	for _, d := range digits {
		if !found0 && len(d) == 6 {
			nums[0] = d
			found0 = true
			continue
		}
		rest = append(rest, d)
	}
	digits = rest
	rest = make([]string, 0)

	// 2
	found2 := false
	for _, d := range digits {
		if !found2 && len(d) == 5 && !containedIn(d, nums[9]) {
			nums[2] = d
			found2 = true
			continue
		}
		rest = append(rest, d)
	}
	digits = rest
	rest = make([]string, 0)

	// 5, 3
	if containedIn(digits[0], nums[6]) {
		nums[5] = digits[0]
		nums[3] = digits[1]
	} else {
		nums[5] = digits[1]
		nums[3] = digits[0]
	}

	return nums
}

type ByRune []rune

func (r ByRune) Len() int {
	return len(r)
}
func (r ByRune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r ByRune) Less(i, j int) bool {
	return r[i] < r[j]
}

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortString(s string) string {
	var r ByRune = stringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

func sortDigits(digits []string) []string {
	for i := 0; i < len(digits); i++ {
		digits[i] = sortString(digits[i])
	}
	return digits
}

func main() {
	signals, outputs, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for i := 0; i < len(signals); i++ {
		digits := sortDigits(makeDigits(signals[i]))

		m := make(map[string]int)
		for i, d := range digits {
			m[d] = i
		}

		tmp := ""
		for _, d := range outputs[i] {
			digit := sortString(d)
			tmp = fmt.Sprintf("%s%d", tmp, m[digit])
		}

		num, err := strconv.Atoi(tmp)
		if err != nil {
			panic(err)
		}

		sum += num
	}
	fmt.Println(sum)
}
