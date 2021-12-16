package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadInput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	number := scanner.Text()
	return number, nil
}

func convertHexToInt64(hex string) int64 {
	hex2int, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		panic(err)
	}
	return hex2int
}

func convertInt64ToBinary(num int64) string {
	return fmt.Sprintf("%04b", num)
}

func convertBinaryToInt64(binary string) int64 {
	num, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func doOperator(opType int64, arr []int64) (int64, error) {
	if opType == 0 {
		sum := int64(0)
		for i := 0; i < len(arr); i++ {
			sum += arr[i]
		}
		return sum, nil
	}
	if opType == 1 {
		prod := arr[0]
		for i := 1; i < len(arr); i++ {
			prod *= arr[i]
		}
		return prod, nil
	}
	if opType == 2 {
		min := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] < min {
				min = arr[i]
			}
		}
		return min, nil
	}
	if opType == 3 {
		max := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max, nil
	}
	if opType == 5 {
		if arr[0] > arr[1] {
			return 1, nil
		}
		return 0, nil
	}
	if opType == 6 {
		if arr[0] < arr[1] {
			return 1, nil
		}
		return 0, nil
	}
	if opType == 7 {
		if arr[0] == arr[1] {
			return 1, nil
		}
		return 0, nil
	}
	return -1, fmt.Errorf("unsupported operation %d", opType)
}

func solve(packet string) (int64, string, error) {
	pType := convertBinaryToInt64(packet[3:6])

	if pType == 4 {
		data := packet[6:]
		seq := ""
		notDone := true
		var i int
		for i = 0; (i < len(data)) && notDone; i += 5 {
			seqId := string(data[i])
			if seqId == "0" {
				notDone = false
			}
			seq = fmt.Sprintf("%s%s", seq, data[i+1:i+5])
		}
		return convertBinaryToInt64(seq), data[i:], nil
	}

	// operator
	lengthTypeId := string(packet[6])
	if lengthTypeId == "0" {
		length := convertBinaryToInt64(packet[7:22])
		data := packet[22:]
		readBits := int64(0)
		sub := make([]int64, 0)
		for readBits < length {
			value, newData, err := solve(data)
			if err != nil {
				return -1, "", err
			}
			sub = append(sub, value)
			readBits += int64(len(data) - len(newData))
			data = newData
		}
		value, err := doOperator(pType, sub)
		if err != nil {
			return -1, "", err
		}
		return value, data, nil
	} else if lengthTypeId == "1" {
		count := convertBinaryToInt64(packet[7:18])
		data := packet[18:]
		sub := make([]int64, 0)
		var value int64
		var err error
		for i := int64(0); i < count; i++ {
			value, data, err = solve(data)
			if err != nil {
				return -1, "", err
			}
			sub = append(sub, value)
		}
		value, err = doOperator(pType, sub)
		if err != nil {
			return -1, "", err
		}
		return value, data, nil
	}
	return -1, "", fmt.Errorf("unsupported length type %s", lengthTypeId)
}

func main() {
	inputHex, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}
	number := ""
	for _, r := range inputHex {
		binary := convertInt64ToBinary(convertHexToInt64(string(r)))
		number = fmt.Sprintf("%s%s", number, binary)
	}
	fmt.Println(solve(number))
}
