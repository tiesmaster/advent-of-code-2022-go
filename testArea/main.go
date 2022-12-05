package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	instruction := parseInstruction("move 4 from 9 to 1")

	fmt.Println(instruction)
}

func parseInstruction(text string) instruction {
	r, _ := regexp.Compile(`move (\d+) from (\d) to (\d)`)
	matches := r.FindStringSubmatch(text)

	quantity, _ := strconv.Atoi(matches[1])
	source, _ := strconv.Atoi(matches[2])
	destination, _ := strconv.Atoi(matches[3])

	return instruction{
		quantity:    quantity,
		source:      source,
		destination: destination,
	}
}

type instruction struct {
	quantity    int
	source      int
	destination int
}
