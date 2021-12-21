package main

import (
	"fmt"
)

func part1() {
	size := 10
	players := []int{0, 0}
	position := []int{6, 9}
	turn := 0
	counter := 1
	loosing := 0
	for {
		position[turn] = ((position[turn] + (counter + (counter + 1) + (counter + 2)) - 1) % size) + 1
		players[turn] += position[turn]
		counter += 3
		if players[turn] >= 1000 {
			loosing = (turn + 1) % 2
			break
		}
		turn = (turn + 1) % 2
	}
	fmt.Printf("%d * %d = %d\n", counter-1, players[loosing], (counter-1)*players[loosing])
}

type Universe struct {
	position1, position2, sum1, sum2 int
}

func part2() {
	// count all possible dice roll
	// example:
	// 3 sum on 3 rolls can be only (1,1,1)=1;
	// 4 sum on 3 rolls can be = (1,1,2),(1,2,1),(2,1,1) = 3; etc...
	rolls := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls[i+j+k]++
			}
		}
	}

	// at the beginning there is only one universe
	universes := make(map[Universe]int)
	universes[Universe{
		position1: 6, // player1 starting position
		position2: 9, // player2 starting position
	}] = 1 // first universe

	turn := 0
	p1Wins := 0
	p2Wins := 0

	// as long as there are universes to explore
	for len(universes) > 0 {
		newUniverses := make(map[Universe]int)
		for universe, universesNum := range universes {
			if universe.sum1 >= 21 {
				// player1 wins, no need to spawn new universes
				p1Wins += universesNum
			} else if universe.sum2 >= 21 {
				// player2 wins, no need to spawn new universes
				p2Wins += universesNum
			} else {
				for rollSum := range rolls {
					newUniverse := Universe{
						position1: universe.position1,
						position2: universe.position2,
						sum1:      universe.sum1,
						sum2:      universe.sum2,
					}
					if turn == 0 {
						newUniverse.position1 = (newUniverse.position1+rollSum-1)%10 + 1
						newUniverse.sum1 += newUniverse.position1
					} else {
						newUniverse.position2 = (newUniverse.position2+rollSum-1)%10 + 1
						newUniverse.sum2 += newUniverse.position2
					}
					// Number rolls[rollSum] tells us how many dice combinations produce the same rollSum number.
					// In other words, we don't go through all combinations, but just to account them into a final count.
					// Since there were universesNum universes before we multiply rolls[rollSum] * universesNum.
					newUniverses[newUniverse] += rolls[rollSum] * universesNum
				}
			}
		}
		turn = (turn + 1) % 2
		universes = newUniverses
	}

	fmt.Printf("player1: %d\n", p1Wins)
	fmt.Printf("player2: %d\n", p2Wins)
	if p1Wins > p2Wins {
		fmt.Println("winner player1")
	} else {
		fmt.Println("winner player2")
	}
}

func main() {
	fmt.Println("Part1:")
	part1()
	fmt.Println("Part2:")
	part2()
}
