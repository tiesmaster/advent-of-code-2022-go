package day15

import (
	"strconv"
	"strings"
)

func Step01(report string) int {
	readings := parseReport(report)
	targetRow := 10
	return calculateOccupiedPositions(readings, targetRow)
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
type rowCoverage struct{ start, end int }

func calculateOccupiedPositions(readings []reading, targetRow int) int {
	coverages := make([]rowCoverage, 0)
	for _, reading := range readings {
		if b, coverage := calculateCoverage(reading, targetRow); b {
			coverages = append(coverages, coverage)
		}
	}
	panic("unimplemented")
}

func calculateCoverage(reading reading, targetRow int) (bool, rowCoverage) {
	distanceTotal := manhattenDistance(reading.sensor, reading.beacon)
	distanceY := distanceToRow(reading.sensor, targetRow)
	halfCoverage := max(0, distanceTotal - distanceY)
	if halfCoverage > 0 {
		sensorX := reading.sensor.x
		return true, rowCoverage{sensorX - halfCoverage, sensorX + halfCoverage}
	} else {
		return false, rowCoverage{0, 0}
	}
}

func manhattenDistance(a, b coordinate) int {
	return distanceX(a, b) + distanceY(a, b)
}

func distanceX(a, b coordinate) int {
	return abs(a.x - b.x)
}

func distanceY(a, b coordinate) int {
	return abs(a.y - b.y)
}

func distanceToRow(coordinate coordinate, targetRow int) int {
	return abs(targetRow - coordinate.y)
}

func toInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}
