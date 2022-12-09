package day09

import (
	"strconv"
	"strings"
)

func Step01(motionsText string) int {
	motions := parseMotions(motionsText)
	state := newStateStep01()
	simulateMotionsStep01(motions, &state)
	return len(state.allTailPositions)
}

func Step02(motionsText string) int {
	motions := parseMotions(motionsText)
	state := newStateStep02()
	simulateMotionsStep02(motions, &state)
	return len(state.allTailPositions)
}

type Direction int

const (
	left  Direction = 0
	right Direction = 1
	up    Direction = 2
	down  Direction = 3
)

type motion struct {
	direction Direction
	steps     int
}

type coordinate struct {
	x int
	y int
}

type StateStep01 struct {
	head             coordinate
	tail             coordinate
	allTailPositions map[coordinate]bool
}

type StateStep02 struct {
	knots            [10]coordinate
	allTailPositions map[coordinate]bool
}

func parseMotions(motionsText string) []motion {
	motions := make([]motion, 0)
	for _, line := range strings.Split(motionsText, "\n") {
		motions = append(motions, parseMotion(line))
	}

	return motions
}

func parseMotion(line string) motion {
	var direction Direction
	steps, _ := strconv.Atoi(line[2:])

	switch line[0] {
	case 'L':
		direction = left
	case 'R':
		direction = right
	case 'U':
		direction = up
	case 'D':
		direction = down
	}

	return motion{direction, steps}
}

func newStateStep01() StateStep01 {
	startPosition := coordinate{0, 0}
	return StateStep01{
		head: startPosition,
		tail: startPosition,
		allTailPositions: map[coordinate]bool{
			startPosition: true,
		},
	}
}

func newStateStep02() StateStep02 {
	startPosition := coordinate{0, 0}
	state := StateStep02{
		allTailPositions: map[coordinate]bool{
			startPosition: true,
		},
	}

	for i := 0; i < len(state.knots); i++ {
		state.knots[i] = startPosition
	}

	return state
}

func simulateMotionsStep01(motions []motion, state *StateStep01) {
	for _, motion := range motions {
		for i := 0; i < motion.steps; i++ {
			state.head.move(motion.direction)

			if !state.head.isAdjacent(state.tail) {
				state.moveTailTowardsHead()
			}
		}
	}
}

func simulateMotionsStep02(motions []motion, state *StateStep02) {
	for _, motion := range motions {
		for i := 0; i < motion.steps; i++ {
			state.knots[0].move(motion.direction)

			for i := 1; i < len(state.knots); i++ {
				curKnot := &state.knots[i]
				prevKnot := state.knots[i-1]
				if !curKnot.isAdjacent(prevKnot) {
					curKnot.moveTowards(prevKnot)
				}
			}
			state.allTailPositions[state.knots[9]] = true
		}
	}
}

func (coordinate *coordinate) move(direction Direction) {
	switch direction {
	case left:
		coordinate.x--
	case right:
		coordinate.x++
	case down:
		coordinate.y--
	case up:
		coordinate.y++
	}
}

func (state *StateStep01) moveTailTowardsHead() {
	state.tail.moveTowards(state.head)
	state.allTailPositions[state.tail] = true

}

func (a *coordinate) moveTowards(b coordinate) {
	if distance(a.x, b.x) > 1 {
		a.x = calculateNewPosition(a.x, b.x)
		a.y = b.y
	} else {
		a.x = b.x
		a.y = calculateNewPosition(a.y, b.y)
	}
}

func calculateNewPosition(a, b int) int {
	if a > b {
		return a - 1
	} else {
		return a + 1
	}
}

func (a *coordinate) isAdjacent(b coordinate) bool {
	return distance(a.x, b.x) < 2 && distance(a.y, b.y) < 2
}

func distance(i, j int) int {
	return abs(i - j)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}
