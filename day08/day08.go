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
	rowCount, columnCount := len(mapData[0]), len(mapData)

	visMap := make([][]bool, columnCount)

	// initialize visMap
	for i := 0; i < len(visMap); i++ {
		visMap[i] = make([]bool, len(mapData[0]))
	}

	// set edges to visible
	for i := 0; i < rowCount; i++ {
		if i == 0 || i == rowCount-1 {
			for j := 0; j < columnCount; j++ {
				visMap[i][j] = true
			}
		} else {
			visMap[i][0] = true
			visMap[i][columnCount-1] = true
		}
	}

	// map out visibility
	for i := 1; i < rowCount; i++ {
		for j := 1; j < columnCount; j++ {
			currentTree := mapData[i][j]

			// x-axis search [leftside]
			for x := 0; x < j; x++ {
				// there is a tree that block line of sight
				if mapData[i][x] >= currentTree {
					break
				}

				// nothing is blocking, so we're visible
				visMap[i][j] = true
			}

			// x-axis search [rightside]
			for x := j+1; x < columnCount; x++ {
				// there is a tree that block line of sight
				if mapData[i][x] >= currentTree {
					break
				}

				// nothing is blocking, so we're visible
				visMap[i][j] = true
			}

			// y-axis search [topside]
			for y := 0; y < j; y++ {
				// there is a tree that block line of sight
				if mapData[y][j] >= currentTree {
					break
				}

				// nothing is blocking, so we're visible
				visMap[i][j] = true
			}

			// y-axis search [bottomside]
			for y := j+1; y < rowCount; y++ {
				// there is a tree that block line of sight
				if mapData[y][j] >= currentTree {
					break
				}

				// nothing is blocking, so we're visible
				visMap[i][j] = true
			}
		}
	}

	// count visible trees
	count := 0
	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			if visMap[i][j] {
				count++
			}
		}
	}

	return count
}
