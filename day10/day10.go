package day10

import (
	"strconv"
	"strings"
)

func Step01(instructionsText string) int {
	instructions := parseInstructions(instructionsText)

	var registerX, sumSignalStrength int

	// var currInstr *instruction
	registerX = 1

	for i, currInstr := range instructions {

		cycle := i+1
		if cycle%40 == 20 {
			sumSignalStrength += registerX * cycle
		}

		if currInstr.operation == addx {
			registerX += currInstr.operarant
		}
	}

	return sumSignalStrength
}

func Step02(instructionsText string) string {
	instructions := parseInstructions(instructionsText)

	var videoBuffer videoBuffer

	var registerX, cycle int

	var currInstr *instruction
	registerX = 1

	for len(instructions) > 0 {
		// fetch next instruction, if needed
		if currInstr == nil {
			currInstr = &instructions[0]
			instructions = instructions[1:]
		}

		videoBuffer.drawPixel(cycle, registerX)

		// start new cycle
		cycle++

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

	return videoBuffer.renderToString()
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
		instructions = append(instructions, parseLine(line)...)
	}
	return instructions
}

func parseLine(line string) []instruction {
	fields := strings.Fields(line)
	switch fields[0] {
	case "addx":
		return addxInstruction(fields[1])
	case "noop":
		return noopInstruction()
	}

	panic("cannot reach")
}

func addxInstruction(operantText string) []instruction {
	V, _ := strconv.Atoi(operantText)

	return append(noopInstruction(), instruction{
		operation: addx,
		operarant: V,
		cycles: 1,
	})
}

func noopInstruction() []instruction {
	return []instruction{
		{
			operation: noop,
			cycles: 1,
		},
	}
}

const (
	numRows    = 6
	numColumns = 40
)

type coordinate struct{ row, col int }
type videoBuffer [numRows][numColumns]bool

func (videoBuffer *videoBuffer) drawPixel(cycle, registerX int) {
	c := toCoordinate(cycle)
	videoBuffer[c.row][c.col] = c.spriteOverlaps(registerX)
}

func toCoordinate(cycle int) coordinate {
	return coordinate{
		row: cycle / numColumns,
		col: cycle % numColumns,
	}
}

func (c coordinate) spriteOverlaps(spriteCol int) bool {
	var low, high int
	if spriteCol < 1 {
		low, high = 0, 2
	} else if spriteCol > numColumns-2 {
		low, high = 37, 39
	} else {
		low, high = spriteCol-1, spriteCol+1
	}

	return low <= c.col && c.col <= high
}

func (videoBuffer *videoBuffer) renderToString() string {
	var lines [numRows]string
	for row := 0; row < numRows; row++ {
		var line [numColumns]byte
		for col := 0; col < numColumns; col++ {
			line[col] = renderPixel(videoBuffer[row][col])
		}
		lines[row] = string(line[:])
	}

	return strings.Join(lines[:], "\n")
}

func renderPixel(b bool) byte {
	if b {
		return '#'
	} else {
		return '.'
	}
}
