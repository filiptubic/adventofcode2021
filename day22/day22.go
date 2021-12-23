package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	from, to int
}

func (r Range) size() int {
	return r.to - r.from + 1
}

type Cube struct {
	action  string
	x, y, z Range
}

func (c Cube) size() int {
	return c.x.size() * c.y.size() * c.z.size()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c Cube) intersect(other Cube) (bool, Cube) {
	xFrom := max(c.x.from, other.x.from)
	xTo := min(c.x.to, other.x.to)
	yFrom := max(c.y.from, other.y.from)
	yTo := min(c.y.to, other.y.to)
	zFrom := max(c.z.from, other.z.from)
	zTo := min(c.z.to, other.z.to)
	if xFrom <= xTo && yFrom <= yTo && zFrom <= zTo {
		return true, Cube{
			action: c.action,
			x:      Range{xFrom, xTo},
			y:      Range{yFrom, yTo},
			z:      Range{zFrom, zTo},
		}
	}
	return false, Cube{}
}

func (c Cube) subtract(other Cube) []Cube {
	ok, other := c.intersect(other)
	if !ok {
		return []Cube{c}
	}

	newCubes := make([]Cube, 0)

	if c.x.from < other.x.from {
		newCubes = append(newCubes, Cube{
			action: c.action,
			x:      Range{c.x.from, other.x.from - 1},
			y:      Range{c.y.from, c.y.to},
			z:      Range{c.z.from, c.z.to},
		})
	}

	if c.x.to > other.x.to {
		newCubes = append(newCubes, Cube{
			action: c.action,
			x:      Range{other.x.to + 1, c.x.to},
			y:      Range{c.y.from, c.y.to},
			z:      Range{c.z.from, c.z.to},
		})
	}

	if c.y.from < other.y.from {
		newCubes = append(newCubes, Cube{
			action: c.action,
			x:      Range{other.x.from, other.x.to},
			y:      Range{c.y.from, other.y.from - 1},
			z:      Range{c.z.from, c.z.to},
		})
	}

	if c.y.to > other.y.to {
		newCubes = append(newCubes, Cube{
			action: c.action,
			x:      Range{other.x.from, other.x.to},
			y:      Range{other.y.to + 1, c.y.to},
			z:      Range{c.z.from, c.z.to},
		})
	}

	if c.z.from < other.z.from {
		newCubes = append(newCubes, Cube{
			action: c.action,
			x:      Range{other.x.from, other.x.to},
			y:      Range{other.y.from, other.y.to},
			z:      Range{c.z.from, other.z.from - 1},
		})
	}

	if c.z.to > other.z.to {
		newCubes = append(newCubes, Cube{
			action: c.action,
			x:      Range{other.x.from, other.x.to},
			y:      Range{other.y.from, other.y.to},
			z:      Range{other.z.to + 1, c.z.to},
		})
	}
	return newCubes
}

func loadInput() []Cube {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	steps := make([]Cube, 0)
	re := regexp.MustCompile("(?P<action>on|off) x=(?P<x>-?[0-9]+..-?[0-9]+),y=(?P<y>-?[0-9]+..-?[0-9]+),z=(?P<z>-?[0-9]+..-?[0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			parts := re.FindStringSubmatch(line)
			action := parts[1]

			xRangeString := strings.Split(parts[2], "..")
			xFrom, err := strconv.Atoi(xRangeString[0])
			if err != nil {
				panic(err)
			}
			xTo, err := strconv.Atoi(xRangeString[1])
			if err != nil {
				panic(err)
			}

			yRangeString := strings.Split(parts[3], "..")
			yFrom, err := strconv.Atoi(yRangeString[0])
			if err != nil {
				panic(err)
			}
			yTo, err := strconv.Atoi(yRangeString[1])
			if err != nil {
				panic(err)
			}

			zRangeString := strings.Split(parts[4], "..")
			zFrom, err := strconv.Atoi(zRangeString[0])
			if err != nil {
				panic(err)
			}
			zTo, err := strconv.Atoi(zRangeString[1])
			if err != nil {
				panic(err)
			}

			steps = append(steps, Cube{
				action: action,
				x:      Range{xFrom, xTo},
				y:      Range{yFrom, yTo},
				z:      Range{zFrom, zTo},
			})
		}
	}
	return steps
}

func count(cubes []Cube, bounds *Cube) int {
	cubesMap := make(map[Cube]struct{})
	for _, newCube := range cubes {
		newCubes := make(map[Cube]struct{})
		if bounds != nil {
			var ok bool
			ok, newCube = newCube.intersect(*bounds)
			if !ok {
				continue
			}
		}

		for oldCube := range cubesMap {
			diffCubes := oldCube.subtract(newCube)
			for _, dc := range diffCubes {
				newCubes[dc] = struct{}{}
			}
		}
		cubesMap = newCubes
		cubesMap[newCube] = struct{}{}
	}

	size := 0
	for c := range cubesMap {
		if c.action == "on" {
			size += c.size()
		}
	}
	return size
}

func main() {
	fmt.Printf("Part1: %d\n", count(loadInput(), &Cube{
		x: Range{-50, 50},
		y: Range{-50, 50},
		z: Range{-50, 50},
	}))
	fmt.Printf("Part2: %d\n", count(loadInput(), nil))
}
