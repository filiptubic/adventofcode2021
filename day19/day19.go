package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	coordsPosition = []struct{ first, second, third int }{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}
	coordsInvert = []struct{ first, second, third int }{
		{1, 1, 1},
		{1, 1, -1},
		{1, -1, 1},
		{1, -1, -1},
		{-1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{-1, -1, -1},
	}
)

type Point struct {
	x, y, z int
}

func (p Point) manhattan(p2 Point) int {
	return int(math.Abs(float64(p.x-p2.x)) + math.Abs(float64(p.y-p2.y)) + math.Abs(float64(p.z-p2.z)))
}

type Tuple2 struct {
	i, j int
}

func loadScans(path string) [][]Point {
	re := regexp.MustCompile("--- scanner [0-9]+ ---")
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scans := make([][]Point, 0)
	scan := make([]Point, 0)
	scanner.Scan()

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if re.MatchString(scanner.Text()) {
			scans = append(scans, scan)
			scan = make([]Point, 0)
			continue
		}
		parts := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		scan = append(scan, Point{x, y, z})
	}
	scans = append(scans, scan)

	return scans
}

func findMatch(scanner1Data, scanner2Data []Point) (bool, []Point, Point) {
	inScanner1 := make(map[Point]struct{})
	for _, x := range scanner1Data {
		inScanner1[x] = struct{}{}
	}

	for _, position := range coordsPosition {
		for _, invert := range coordsInvert {

			scanner2NewPoints := make([]Point, len(scanner2Data))
			for i := 0; i < len(scanner2Data); i++ {
				coords := []int{scanner2Data[i].x, scanner2Data[i].y, scanner2Data[i].z}
				scanner2NewPoints[i] = Point{
					x: coords[position.first] * invert.first,
					y: coords[position.second] * invert.second,
					z: coords[position.third] * invert.third,
				}
			}

			for _, pointScanner1 := range scanner1Data {
				for _, pointScanner2 := range scanner2NewPoints {
					scanner2Position := Point{
						x: pointScanner2.x - pointScanner1.x,
						y: pointScanner2.y - pointScanner1.y,
						z: pointScanner2.z - pointScanner1.z,
					}

					matches := 0
					relativePositions := make([]Point, 0)
					for _, otherPointScanner2 := range scanner2NewPoints {
						positionRelativeToScanner1 := Point{
							x: otherPointScanner2.x - scanner2Position.x,
							y: otherPointScanner2.y - scanner2Position.y,
							z: otherPointScanner2.z - scanner2Position.z,
						}
						if _, ok := inScanner1[positionRelativeToScanner1]; ok {
							matches++
						}
						relativePositions = append(relativePositions, positionRelativeToScanner1)
					}

					if matches >= 12 {
						return true, relativePositions, scanner2Position
					}
				}
			}
		}
	}
	return false, nil, Point{}
}

func main() {
	scans := loadScans("input.txt")
	scanners := make([]Point, 0)
	skip := make(map[Tuple2]struct{})
	found := make(map[int][]Point)
	found[0] = scans[0]
	allFound := make(map[Point]struct{})
	for _, p := range scans[0] {
		allFound[p] = struct{}{}
	}

	for len(found) < len(scans) {
		for i, scan := range scans {
			if _, ok := found[i]; ok {
				continue
			}

			for j := range found {
				if _, ok := skip[Tuple2{i, j}]; ok {
					continue
				}

				if ok, positions, scanner := findMatch(found[j], scan); ok {
					scanners = append(scanners, scanner)
					found[i] = positions
					for _, m := range positions {
						allFound[m] = struct{}{}
					}
					break
				}
				skip[Tuple2{i, j}] = struct{}{}
			}
		}
	}

	fmt.Printf("Part1: %d\n", len(allFound))
	max := -1
	for _, scan1 := range scanners {
		for _, scan2 := range scanners {
			distance := scan1.manhattan(scan2)
			if distance > max {
				max = distance
			}
		}
	}
	fmt.Printf("Part2: %d\n", max)
}
