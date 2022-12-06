package day05

import (
	"fmt"
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

func PerformRearrangementProcedureButWithCrateMover9001(drawing string) string {
	parts := strings.Split(drawing, "\n\n")

	state := parseStartingState(parts[0])
	instructions := parseInstructions(parts[1])

	for _, instr := range instructions {
		executeInstructionViaCrateMover9001(state, instr)
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
	parts := strings.Split(stateText, "\n")

	totalLineCount := len(parts)
	totalCratesCount := parseCrateNumbers(parts[totalLineCount-1])

	crates := parseCrates(parts[:totalLineCount-1], totalCratesCount)

	fmt.Println(crates)

	return State{crates}
}

func parseCrateNumbers(crateIndexes string) int {
	crateNumbers := strings.Split(crateIndexes, "  ")

	s := crateNumbers[len(crateNumbers)-1]
	s = strings.Trim(s, " ")
	lastCrateNumber, _ := strconv.Atoi(s)

	return lastCrateNumber
}

func parseCrates(crateLines []string, totalCratesCount int) []stack {
	crates := make([]stack, totalCratesCount)

	for i := len(crateLines) - 1; i >= 0; i-- {
		line := crateLines[i]
		for j := 0; j < totalCratesCount; j++ {
			index := 4*j + 1
			value := line[index]
			if value != ' ' {
				crates[j] = append(crates[j], value)
			}
		}
	}

	return crates

}

func parseInstructions(instructionText string) []instruction {
	parts := strings.Split(instructionText, "\n")

	instructions := make([]instruction, len(parts))
	for i, text := range parts {
		instructions[i] = parseInstruction(text)
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
		quantity: quantity,
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

func executeInstructionViaCrateMover9001(state State, instr instruction) {
	s, d := move(state.crates[instr.source], state.crates[instr.destination], instr.quantity)
	state.crates[instr.source] = s
	state.crates[instr.destination] = d
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

func move(sourceStack stack, destinationStack stack, moveCount int) (stack, stack) {
	totalCountToMove := min(moveCount, len(sourceStack))
	toMove := sourceStack[len(sourceStack)-totalCountToMove:]
	destinationStack = append(destinationStack, toMove...)
	return sourceStack[:len(sourceStack)-totalCountToMove], destinationStack
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
