package main

import (
	"fmt"
	"math"
)

const (
	xMin, xMax = 139, 187
	yMin, yMax = -148, -89
)

type Probe struct {
	yMax                 int
	x, y                 int
	xVelocity, yVelocity int
}

func (p Probe) isInside() bool {
	return (p.x >= xMin) && (p.x <= xMax) && (p.y >= yMin) && (p.y <= yMax)
}

func (p Probe) hasPassed() bool {
	return p.x > xMax || p.y < yMin
}

func checkProbe(p *Probe) bool {
	for {
		if p.isInside() {
			return true
		} else if p.hasPassed() {
			return false
		} else {
			p.x += p.xVelocity
			p.y += p.yVelocity
			if p.xVelocity > 0 {
				p.xVelocity--
			}
			p.yVelocity--
			if p.y > p.yMax {
				p.yMax = p.y
			}
			continue
		}
	}
}

func launchProbes() (int, int) {
	max := 0
	hits := 0
	for x := 0; x < xMax+1; x++ {
		for y := yMin - 1; y < int(math.Abs(float64(yMin)))+1; y++ {
			p := Probe{xVelocity: x, yVelocity: y}
			if checkProbe(&p) {
				if p.yMax > max {
					max = p.yMax
				}
				hits++
			}
		}
	}
	return max, hits
}

func main() {
	fmt.Println(launchProbes())
}
