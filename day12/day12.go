package day12

import (
	"fmt"
	"strings"
)

func Step01(gridData string) int {
	grid, start, end := parseGrid(gridData)

	fmt.Println(grid, start, end)
	// do the algorithm
	panic("unimplemented")
}

type coordinate struct{ x, y int }

func parseGrid(gridData string) (grid [][]int, start, end coordinate) {
	rowCount, columnCount := readDimensions(gridData)

	// initialize grid
	grid = make([][]int, rowCount)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, columnCount)
	}

	start, end = readGrid(gridData, grid)
	return grid, start, end
}

func readDimensions(gridData string) (rowCount, columnCount int) {
	columnCount = strings.Index(gridData, "\n")
	rowCount = strings.Count(gridData, "\n")

	if gridData[len(gridData)-1] != '\n' {
		// account for missing newline at the end of the mapText
		rowCount++
	}

	return
}

func readGrid(gridData string, grid [][]int) (start, end coordinate) {
	lines := strings.Split(gridData, "\n")
	for i, line := range lines {
		for j, r := range line {
			var x int
			switch r {
			case 'S':
				start = coordinate{i, j}
				x = 0
			case 'E':
				end = coordinate{i, j}
				x = 25
			default:
				x = int(r - 'a')
			}
			grid[i][j] = x
		}
	}
	return
}
