package main

import (
	"container/heap"
	"crypto/sha256"
	"fmt"
	"math"
)

type Amphipod int

const (
	Missing Amphipod = 0
	A       Amphipod = 1
	B       Amphipod = 10
	C       Amphipod = 100
	D       Amphipod = 1000
)

var (
	roomIDs = map[Amphipod]int{
		A: 0,
		B: 1,
		C: 2,
		D: 3,
	}
	roomDoorToAmphipod = map[int]Amphipod{
		2: A,
		4: B,
		6: C,
		8: D,
	}
	amphipodToRoomDoor = map[Amphipod]int{
		A: 2,
		B: 4,
		C: 6,
		D: 8,
	}
)

type Room struct {
	amphipodType Amphipod
	size         int
	amphipods    []Amphipod
}

func (r Room) ok() bool {
	if len(r.amphipods) == r.size {
		for _, a := range r.amphipods {
			if a != r.amphipodType {
				return false
			}
		}
		// room is in final state
		return true
	}
	return false
}

func (r Room) canEnter() bool {
	if len(r.amphipods) < r.size {
		for _, a := range r.amphipods {
			if a != r.amphipodType {
				return false
			}
		}
		return true
	}
	return false
}

func (r Room) push(a Amphipod) Room {
	return Room{
		amphipodType: r.amphipodType,
		size:         r.size,
		amphipods:    append([]Amphipod{a}, r.amphipods...),
	}
}

func (r Room) pop() (Amphipod, Room) {
	return r.amphipods[0], Room{
		amphipodType: r.amphipodType,
		size:         r.size,
		amphipods:    r.amphipods[1:],
	}
}

func (r Room) freeSpace() int {
	return r.size - len(r.amphipods)
}

type Hallway struct {
	positions []Amphipod
}

func isDoor(index int) bool {
	if _, ok := roomDoorToAmphipod[index]; ok {
		return true
	}
	return false
}

type Move struct {
	position, distance int
}

func (h Hallway) findMoves(start int) []Move {
	m := make([]Move, 0)

	//left
	for i := start - 1; i >= 0; i-- {
		if isDoor(i) {
			continue
		}

		if h.positions[i] != Missing {
			break
		}
		m = append(m, Move{position: i, distance: int(math.Abs(float64(start - i)))})
	}

	// right
	for i := start + 1; i < len(h.positions); i++ {
		if isDoor(i) {
			continue
		}

		if h.positions[i] != Missing {
			break
		}

		m = append(m, Move{position: i, distance: int(math.Abs(float64(start - i)))})
	}

	return m
}

func (h Hallway) isClear(start, stop int) bool {
	if start < stop {
		for i := start + 1; i <= stop; i++ {
			if h.positions[i] != Missing {
				return false
			}
		}
		return true
	} else {
		for i := stop; i < start; i++ {
			if h.positions[i] != Missing {
				return false
			}
		}
		return true
	}
}

func (h Hallway) update(position int, amphipod Amphipod) Hallway {
	spaces := make([]Amphipod, len(h.positions))
	for i := 0; i < len(spaces); i++ {
		spaces[i] = h.positions[i]
	}
	spaces[position] = amphipod
	return Hallway{
		positions: spaces,
	}
}

type State struct {
	energy  int
	rooms   [4]Room
	hallway Hallway
}

func (s State) ok() bool {
	for _, room := range s.rooms {
		if !room.ok() {
			return false
		}
	}
	// all rooms are in final state
	return true
}

func updateRoom(rooms [4]Room, r Room, index int) [4]Room {
	newRooms := [4]Room{
		rooms[0],
		rooms[1],
		rooms[2],
		rooms[3],
	}
	newRooms[index] = r
	return newRooms
}

type StateHash struct {
	hallway [11]int
	rooms   [4][]int
}

// id uniquely identifies state. This has a huge impact on performance while keeping track of visited states.
func (s State) id() string {
	stateHash := StateHash{}
	for i, a := range s.hallway.positions {
		stateHash.hallway[i] = int(a)
	}
	for i, r := range s.rooms {
		stateHash.rooms[i] = make([]int, len(r.amphipods))
		for j, a := range r.amphipods {
			stateHash.rooms[i][j] = int(a)
		}
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", stateHash)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func minimalEnergy(initialState State) int {
	q := make(PriorityQueue, 0)
	heap.Push(&q, &Item{
		value:    &initialState,
		priority: initialState.energy,
	})
	heap.Init(&q)
	visited := make(map[string]struct{})

	for q.Len() > 0 {
		state := *(heap.Pop(&q).(*Item).value)
		if state.ok() {
			return state.energy
		}
		if _, ok := visited[state.id()]; ok {
			continue
		}
		visited[state.id()] = struct{}{}

		for i, room := range state.rooms {
			if len(room.amphipods) > 0 && !room.ok() {
				amphipod, newRoom := room.pop()

				for _, move := range state.hallway.findMoves(amphipodToRoomDoor[room.amphipodType]) {
					newState := State{
						energy:  state.energy + int(amphipod)*(room.freeSpace()+move.distance+1),
						rooms:   updateRoom(state.rooms, newRoom, i),
						hallway: state.hallway.update(move.position, amphipod),
					}
					if _, ok := visited[newState.id()]; !ok {
						heap.Push(&q, &Item{
							value:    &newState,
							priority: newState.energy,
						})
					}
				}
			}
		}

		for space, amphipod := range state.hallway.positions {
			if amphipod == Missing {
				continue
			}

			door := amphipodToRoomDoor[amphipod]
			roomId := roomIDs[amphipod]
			room := state.rooms[roomId]

			if state.hallway.isClear(space, door) && room.canEnter() {
				newRoom := room.push(amphipod)
				distance := int(math.Abs(float64(door - space)))

				newState := State{
					energy:  state.energy + int(amphipod)*(distance+room.freeSpace()),
					rooms:   updateRoom(state.rooms, newRoom, roomId),
					hallway: state.hallway.update(space, Missing),
				}

				if _, ok := visited[newState.id()]; !ok {
					heap.Push(&q, &Item{
						value:    &newState,
						priority: newState.energy,
					})
				}
			}
		}
	}

	return -1
}

func part1() {
	startHallway := Hallway{positions: make([]Amphipod, 11)}
	for i := 0; i < 11; i++ {
		startHallway.positions[i] = Missing
	}
	startRooms := [4]Room{
		{
			amphipodType: A,
			size:         2,
			amphipods:    []Amphipod{C, C},
		},
		{
			amphipodType: B,
			size:         2,
			amphipods:    []Amphipod{A, A},
		},
		{
			amphipodType: C,
			size:         2,
			amphipods:    []Amphipod{B, D},
		},
		{
			amphipodType: D,
			size:         2,
			amphipods:    []Amphipod{D, B},
		},
	}

	energy := minimalEnergy(State{
		energy:  0,
		hallway: startHallway,
		rooms:   startRooms,
	})

	fmt.Printf("Part1: %d\n", energy)
}

func part2() {
	startHallway := Hallway{positions: make([]Amphipod, 11)}
	for i := 0; i < 11; i++ {
		startHallway.positions[i] = Missing
	}
	startRooms := [4]Room{
		{
			amphipodType: A,
			size:         4,
			amphipods:    []Amphipod{C, D, D, C},
		},
		{
			amphipodType: B,
			size:         4,
			amphipods:    []Amphipod{A, C, B, A},
		},
		{
			amphipodType: C,
			size:         4,
			amphipods:    []Amphipod{B, B, A, D},
		},
		{
			amphipodType: D,
			size:         4,
			amphipods:    []Amphipod{D, A, C, B},
		},
	}

	energy := minimalEnergy(State{
		energy:  0,
		hallway: startHallway,
		rooms:   startRooms,
	})

	fmt.Printf("Part2: %d\n", energy)
}

func main() {
	part1()
	part2()
}
