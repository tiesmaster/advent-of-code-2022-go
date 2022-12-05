package day05

import (
	// "strconv"
	"strings"
)

func PerformRearrangementProcedure(drawing string) string {
	parts := strings.Split(drawing, "\n\n")

	state := parseState(parts[0])
	instructions := parseInstructions(parts[1])

	for _, instr := range instructions {
		executeInstruction(state, instr)
	}

	return topCrates(state)
}

type State []crateStack

type crateStack []string

type instruction struct {
	quantity    int
	source      int
	destination int
}

func parseState(s string) State {
	panic("unimplemented")
}

func parseInstructions(instructionText string) []instruction {
	panic("unimplemented")
}

func executeInstruction(state State, instr instruction) {
	panic("unimplemented")
}

func topCrates(state State) string {
	panic("unimplemented")
}
