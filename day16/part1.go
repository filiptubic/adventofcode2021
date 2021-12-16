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

func solve(packet string) (int64, string, error) {
	pType := convertBinaryToInt64(packet[3:6])
	versionSum := convertBinaryToInt64(packet[:3])

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
		return versionSum, data[i:], nil
	}

	// operator
	lengthTypeId := string(packet[6])
	if lengthTypeId == "0" {
		length := convertBinaryToInt64(packet[7:22])
		data := packet[22:]
		readBits := int64(0)
		for readBits < length {
			version, newData, err := solve(data)
			if err != nil {
				return -1, "", err
			}
			readBits += int64(len(data) - len(newData))
			data = newData
			versionSum += version
		}
		return versionSum, data, nil
	} else if lengthTypeId == "1" {
		count := convertBinaryToInt64(packet[7:18])
		data := packet[18:]
		var version int64
		var err error
		for i := int64(0); i < count; i++ {
			version, data, err = solve(data)
			if err != nil {
				return -1, "", err
			}
			versionSum += version
		}
		return versionSum, data, nil
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
