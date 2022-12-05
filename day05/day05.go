package day05

import (
	// "strconv"
	"regexp"
	"strconv"
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
	crates := []stack{
		{90, 78},
		{77, 67, 68},
		{80},
	}

	return State{
		crates: crates,
	}
}

func parseInstructions(instructionText string) []instruction {
	parts := strings.Split(instructionText, "\n")

	instructions := make([]instruction, 0)
	for _, text := range parts {
		instructions = append(instructions, parseInstruction(text))
	}

	return instructions
}

func parseInstruction(text string) instruction {
	r, _ := regexp.Compile(`move (\d+) from (\d) to (\d)`)
	matches := r.FindStringSubmatch(text)

	quantity, _ := strconv.Atoi(matches[1])
	source, _ := strconv.Atoi(matches[2])
	destination, _ := strconv.Atoi(matches[3])

	return instruction{
		quantity:    quantity,
		// the instructions are one-based, instead of zero-based
		source:      source - 1,
		destination: destination - 1,
	}
}

func executeInstruction(state State, instr instruction) {
	var b byte
	for i := 0; i < instr.quantity; i++ {
		state.crates[instr.source], b = pop(state.crates[instr.source])
		state.crates[instr.destination] = push(state.crates[instr.destination], b)
	}
}

func topCrates(state State) string {
	bytes := make([]byte, 0)
	for i := 0; i < len(state.crates); i++ {
		crate := state.crates[i]
		if len(crate) > 0 {
			bytes = append(bytes, crate[len(crate)-1])
		}
	}

	// TODO: make use of builder
	// https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/

	return string(bytes)
}

func push(st stack, item byte) stack {
	return append(st, item)
}

func pop(st stack) (stack, byte) {
	n := len(st) - 1
	ret := st[n]

	st = st[:n]

	return st, ret
}
