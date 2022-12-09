package day09

import (
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/cloud/location"
)

func Step01(motionsText string) int {
	motions := parseMotions(motionsText)
	state := newState()
	simulateMotions(motions, &state)
	return len(state.previousTailPositions)
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

type State struct {
	head                  coordinate
	tail                  coordinate
	previousTailPositions map[coordinate]bool
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

func newState() State {
	return State{
		head:                  coordinate{0, 0},
		tail:                  coordinate{0, 0},
		previousTailPositions: make(map[coordinate]bool),
	}
}

func simulateMotions(motions []motion, state *State) {
	panic("unimplemented")
}
