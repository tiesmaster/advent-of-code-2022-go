package day08

import (
	"strconv"
	"strings"
)

func Step01(mapText string) int {
	mapData := parseMap(mapText)
	return calculateVisibleTrees(mapData)
}

func parseMap(mapText string) [][]int {
	rowCount, columnCount := readDimensions(mapText)
	m := make([][]int, columnCount)

	// initialize map
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, rowCount)
	}

	readMap(mapText, m)

	return m
}

func readDimensions(mapText string) (rowCount, columnCount int) {
	columnCount = strings.Index(mapText, "\n")
	rowCount = strings.Count(mapText, "\n")

	if mapText[len(mapText)-1] != '\n' {
		// account for missing newline at the end of the mapText
		rowCount++
	}

	return
}

func readMap(mapText string, mapData [][]int) {
	lines := strings.Split(mapText, "\n")
	for i, line := range lines {
		for j, r := range line {
			x, _ := strconv.Atoi(string(r))
			mapData[i][j] = x
		}
	}
}

func calculateVisibleTrees(mapData [][]int) int {
	panic("unimplemented")
}
