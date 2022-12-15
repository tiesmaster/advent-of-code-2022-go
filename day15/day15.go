package day15

import (
	"strconv"
	"strings"
)

func Step01(report string) int {
	readings := parseReport(report)
	targetRow := 10
	return calculateCoverage(readings, targetRow)
}

func parseReport(report string) []reading {
	lines := strings.Split(report, "\n")
	readings := make([]reading, len(lines))
	for i, line := range lines {
		readings[i] = parseLine(line)
	}
	return readings
}

func parseLine(line string) reading {
	parts := strings.Split(line, ": ")
	return reading{parsePart(parts[0]), parsePart(parts[1])}
}

func parsePart(part string) coordinate {
	const atString = " at "
	index := strings.Index(part, atString)
	coordinateString := part[index+len(atString):]
	return parseCoordinate(coordinateString)
}

func parseCoordinate(coordinateString string) coordinate {
	parts := strings.Split(coordinateString, ", ")
	x := toInt(parts[0][2:])
	y := toInt(parts[1][2:])
	return coordinate{x, y}
}

type reading struct{ sensor, beacon coordinate }
type coordinate struct{ x, y int }

func calculateCoverage(readings []reading, targetRow int) int {
	panic("unimplemented")
}

func toInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
