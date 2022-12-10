package day10

import (
	"strconv"
	"strings"
)

func Step01(instructionsText string) int {
	instructions := parseInstructions(instructionsText)

	var registerX, cycle, sumSignalStrength int

	var currInstr *instruction
	registerX = 1

	for len(instructions) > 0 {

		// start new cycle
		cycle++

		if cycle%40 == 20 {
			sumSignalStrength += registerX * cycle
		}

		// fetch next instruction, if needed
		if currInstr == nil {
			currInstr = &instructions[0]
			instructions = instructions[1:]
		}

		// consume cycles, by instruction
		currInstr.cycles--

		// check for completion
		if currInstr.cycles == 0 {
			if currInstr.operation == addx {
				registerX += currInstr.operarant
			}
			currInstr = nil
		}

	}

	return sumSignalStrength
}

type instruction struct {
	operation operation
	operarant int
	cycles    int
}

type operation int

const (
	addx operation = 1
	noop operation = 2
)

func parseInstructions(instructionsText string) []instruction {
	lines := strings.Split(instructionsText, "\n")
	instructions := make([]instruction, 0)
	for _, line := range lines {
		instructions = append(instructions, parseLine(line))
	}
	return instructions
}

func parseLine(line string) instruction {
	fields := strings.Fields(line)
	switch fields[0] {
	case "addx":
		return addxInstruction(fields[1])
	case "noop":
		return noopInstruction()
	}

	panic("cannot reach")
}

func addxInstruction(operantText string) instruction {
	V, _ := strconv.Atoi(operantText)
	return instruction{
		operation: addx,
		operarant: V,
		cycles: 2,
	}
}

func noopInstruction() instruction {
	return instruction{
		operation: noop,
		cycles: 1,
	}
}
