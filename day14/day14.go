package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func Step01(scan string) int {
	paths := parseScan(scan)
	fmt.Println(paths)
	panic("unimplemented")
}

type path []coordinate

type coordinate struct{ x, y int }

func parseScan(scan string) []path {
	lines := strings.Split(scan, "\n")
	paths := make([]path, len(lines))
	for i, line := range lines {
		paths[i] = parseLine(line)
	}
	return paths
}

func parseLine(line string) path {
	parts := strings.Split(line, " -> ")
	coordinates := make([]coordinate, len(parts))
	for i, c := range parts {
		coordinates[i] = parseCoordinate(c)
	}
	return coordinates
}

func parseCoordinate(s string) coordinate {
	tupple := strings.Split(s, ",")
	return coordinate{toInt(tupple[0]), toInt(tupple[1])}
}

func toInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
