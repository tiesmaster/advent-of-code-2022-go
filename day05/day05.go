package day05

import (
	// "strconv"
	"strings"
)

func PerformRearrangementProcedure(drawing string) string {
	parts := strings.Split(drawing, "\n\n")

	state := parseStartingState(parts[0])
	instructions := parseInstructions(parts[1])

	for _, instr := range instructions {
		executeInstruction(state, instr)
	}

	return topCrates(state)
}

type State struct {
	crates []stack
}

type stack []byte

type instruction struct {
	quantity    int
	source      int
	destination int
}

func parseStartingState(stateText string) State {
	panic("unimplemented")
}

func parseInstructions(instructionText string) []instruction {
	panic("unimplemented")
}

func executeInstruction(state State, instr instruction) {
	panic("unimplemented")
}

func topCrates(state State) string {
	s := make([]byte, 0)
	for i := 0; i < len(state.crates); i++ {
		crate := state.crates[i]
		if len(crate) > 0 {
			s = append(s, crate[len(crate)-1])
		}
	}

	// TODO: make use of builder
	// https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/

	return string(s)
}
