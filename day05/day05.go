package day05

import (
	"regexp"
	"strconv"
	"strings"
)

func PerformRearrangementProcedure(drawing string) string {
	state, instructions := parseDrawing(drawing)
	for _, instr := range instructions {
		executeInstruction(state, instr)
	}

	return topCrates(state)
}

func PerformRearrangementProcedureButWithCrateMover9001(drawing string) string {
	state, instructions := parseDrawing(drawing)
	for _, instr := range instructions {
		executeInstructionViaCrateMover9001(state, instr)
	}

	return topCrates(state)
}

type State struct {
	crates []stack
}

type stack struct{
	slice []rune
}

type instruction struct {
	quantity    int
	source      int
	destination int
}

func parseDrawing(drawing string) (State, []instruction) {
	parts := strings.Split(drawing, "\n\n")

	state := parseStartingState(parts[0])
	instructions := parseInstructions(parts[1])
	return state, instructions
}

func parseStartingState(stateText string) State {
	parts := strings.Split(stateText, "\n")

	totalLineCount := len(parts)
	totalCratesCount := parseCrateNumbers(parts[totalLineCount-1])

	crates := parseCrates(parts[:totalLineCount-1], totalCratesCount)

	return State{crates}
}

func parseCrateNumbers(crateIndexes string) int {
	crateNumbers := strings.Fields(crateIndexes)
	lastCrateNumber, _ := strconv.Atoi(crateNumbers[len(crateNumbers)-1])

	return lastCrateNumber
}

func parseCrates(crateLines []string, totalCratesCount int) []stack {
	crates := createCrates(totalCratesCount)

	for i := len(crateLines) - 1; i >= 0; i-- {
		line := crateLines[i]
		for j := 0; j < totalCratesCount; j++ {
			index := 4*j + 1
			value := rune(line[index])
			if value != ' ' {
				crates[j].push(value)
			}
		}
	}

	return crates

}

func createCrates(totalCratesCount int) []stack {
	crates := make([]stack, totalCratesCount)

	for i := 0; i < totalCratesCount; i++ {
		crates[i] = newStack()
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
	for i := 0; i < instr.quantity; i++ {
		state.crates[instr.destination].push(state.crates[instr.source].pop())
	}
}

func executeInstructionViaCrateMover9001(state State, instr instruction) {
	tempStack := newStack()
	for i := 0; i < instr.quantity; i++ {
		tempStack.push(state.crates[instr.source].pop())
	}
	
	for !tempStack.isEmpty() {
		state.crates[instr.destination].push(tempStack.pop())
	}
}

func topCrates(state State) string {
	s := ""
	for i := 0; i < len(state.crates); i++ {
		crate := state.crates[i]
		if !crate.isEmpty() {
			b := crate.pop()
			s = s + string(b)
		}
	}

	return s
}

func newStack() stack {
	return stack{make([]rune, 0)}
}

func (st *stack) push(item rune) {
	st.slice = append(st.slice, item)
}

func (st *stack) pop() rune {
	n := len(st.slice) - 1
	ret := st.slice[n]

	st.slice = st.slice[:n]

	return ret
}

func (st *stack) isEmpty() bool {
	return len(st.slice) == 0
}
