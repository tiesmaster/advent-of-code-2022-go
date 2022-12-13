package day12

import (
	"fmt"
	"math"
	"strings"
)

func Step01(gridData string) int {
	grid, start, end := parseGrid(gridData)

	fmt.Println(grid, start, end)
	// do the algorithm
	panic("unimplemented")

	// path := make([]coordinate, 0)
	// current := start
	// for current != end {
	// 	options := possibleNextSteps(grid, current)
	// 	current = bestStep(options, end)
	// 	path = append(path, current)
	// }

	// return len(path)
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

func distance(a, b coordinate) float64 {

	v := math.Pow(float64(a.x-b.x), 2)
	w := math.Pow(float64(a.y-b.y), 2)
	return math.Abs(math.Sqrt(v + w))
}

func possibleNextSteps(grid [][]int, position coordinate) []coordinate {
	panic("unimplemented")
	// nextSteps := make([]coordinate, 0)

	// left: x - 1, y
	// top:
	// right
	// bottom
	// if grid[coordinate.x-1][coordinate.y] {

	// }
}
