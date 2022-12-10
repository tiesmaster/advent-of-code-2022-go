package day10

import (
	"strconv"
	"strings"
)

func Step01(instructionsText string) int {
	instructions := parseInstructions(instructionsText)

	var registerX, cycle, sumSignalStrength int
	registerX = 1
	for _, instr := range instructions {
		switch instr.operation {
		case addx:
			cycle += 2
			registerX += instr.operarant
		case noop:
			cycle++
		}

		if cycle%40 == 20 {
			sumSignalStrength += registerX * cycle
		}
	}

	return sumSignalStrength
}

type instruction struct {
	operation operation
	operarant int
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
	return instruction{operation: addx, operarant: V}
}

func noopInstruction() instruction {
	return instruction{operation: noop}
}
