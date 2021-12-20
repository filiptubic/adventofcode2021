package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func loadInput(path string) (string, [][]string) {
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	algorithm := scanner.Text()

	image := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		part := make([]string, 0)
		for _, r := range line {
			part = append(part, string(r))
		}
		image = append(image, part)
	}
	return algorithm, image
}

type Tuple2 struct {
	i, j int
}

func (t Tuple2) String() string {
	return fmt.Sprintf("(%d, %d)", t.i, t.j)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func addPixels(image map[Tuple2]string, pixel string) map[Tuple2]string {
	// MxN image
	newImage := make(map[Tuple2]string)
	for k, v := range image {
		newImage[k] = v
	}
	image = newImage

	minN := math.MaxInt64
	minM := math.MaxInt64
	maxN := math.MinInt64
	maxM := math.MinInt64

	for t := range image {
		minM = minInt(minM, t.i)
		minN = minInt(minN, t.j)

		maxM = maxInt(maxM, t.i)
		maxN = maxInt(maxN, t.j)
	}

	for j := minN; j < maxN+1; j++ {
		image[Tuple2{minM - 1, j}] = pixel
		image[Tuple2{maxM + 1, j}] = pixel
	}

	for i := minM; i < maxM+1; i++ {
		image[Tuple2{i, minN - 1}] = pixel
		image[Tuple2{i, maxN + 1}] = pixel
	}

	image[Tuple2{minM - 1, minN - 1}] = pixel
	image[Tuple2{minM - 1, maxN + 1}] = pixel
	image[Tuple2{maxM + 1, minN - 1}] = pixel
	image[Tuple2{maxM + 1, maxN + 1}] = pixel

	return image
}

func newRow(pixel string, t Tuple2, image map[Tuple2]string, directions []Tuple2) []string {
	row := make([]string, 0)
	for _, dir := range directions {
		newT := Tuple2{
			i: t.i + dir.i,
			j: t.j + dir.j,
		}
		if _, ok := image[newT]; !ok {
			row = append(row, pixel)
		} else {
			row = append(row, image[newT])
		}
	}
	return row
}

func enhance(image map[Tuple2]string, algorithm, pixel string) map[Tuple2]string {
	enhanced := make(map[Tuple2]string)

	directionsTop := []Tuple2{{-1, -1}, {-1, 0}, {-1, 1}}
	directionsMid := []Tuple2{{0, -1}, {0, 0}, {0, 1}}
	directionsBottom := []Tuple2{{1, -1}, {1, 0}, {1, 1}}

	keys := make([]Tuple2, 0)
	for k := range image {
		keys = append(keys, k)
	}

	for _, t := range keys {
		pixelsArray := []string{
			strings.Join(newRow(pixel, t, image, directionsTop), ""),
			strings.Join(newRow(pixel, t, image, directionsMid), ""),
			strings.Join(newRow(pixel, t, image, directionsBottom), ""),
		}

		pixels := ""
		for _, p := range pixelsArray {
			pixels = fmt.Sprintf("%s%s", pixels, p)
		}

		bits := ""
		for _, p := range pixels {
			if p == '#' {
				bits = fmt.Sprintf("%s1", bits)
			} else {
				bits = fmt.Sprintf("%s0", bits)
			}
		}

		index, err := strconv.ParseInt(bits, 2, 64)
		if err != nil {
			panic(err)
		}

		enhanced[t] = string(algorithm[index])
	}
	return enhanced
}

func enhanceImage(image map[Tuple2]string, algorithm string, steps int) int {
	for step := 0; step < steps; step++ {
		pixel := "."
		if step%2 != 0 {
			pixel = string(algorithm[0])
		}
		image = addPixels(image, pixel)
		image = enhance(image, algorithm, pixel)
	}
	counter := 0
	for _, v := range image {
		if v == "#" {
			counter++
		}
	}
	return counter
}

func main() {
	algorithm, matrix := loadInput("input.txt")
	image := make(map[Tuple2]string)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			image[Tuple2{i, j}] = matrix[i][j]
		}
	}

	fmt.Printf("Part1: %d\n", enhanceImage(image, algorithm, 2))
	fmt.Printf("Part2: %d\n", enhanceImage(image, algorithm, 50))
}
