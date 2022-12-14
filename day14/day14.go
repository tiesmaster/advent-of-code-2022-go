package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func Step01(scan string) int {
	paths := parseScan(scan)
	startingCoord := coordinate{500, 0}

	gridBoundingBox := getBoundingBox(paths).union(startingCoord.boundingBox())
	grid := makeGrid(gridBoundingBox)
	grid.drawPaths(paths)
	// printGrid(grid)
	return simulateFalling(grid, startingCoord)
}

func Step02(scan string) int {
	paths := parseScan(scan)
	startingCoord := coordinate{500, 0}

	gridBoundingBox := getBoundingBox(paths).union(startingCoord.boundingBox())
	groundFloor := gridBoundingBox.rightTop.moveDown().moveDown()
	gridBoundingBox = gridBoundingBox.union(groundFloor.boundingBox())
	grid := makeGrid(gridBoundingBox)
	grid.drawPaths(paths)
	grid.drawGround(groundFloor)

	printGrid(grid)
	return simulateFallingStep2(grid, startingCoord)
}

func (grid grid) drawPaths(paths []path) {
	for _, p := range paths {
		grid.drawPath(p)
	}
}

func (grid grid) drawPath(path path) {
	for i := 0; i < len(path)-1; i++ {
		grid.drawLine(Line{path[i], path[i+1]})
	}
}

func (grid grid) drawLine(line Line) {
	line = normalize(line)
	// fmt.Println("Drawing line: ", line)
	switch {
	case line[0].y == line[1].y:
		// horizontal line
		y := line[0].y
		for x := line[0].x; x < line[1].x+1; x++ {
			grid.bitmap[x][y] = true
		}
	case line[0].x == line[1].x:
		// vertical line
		x := line[0].x
		for y := line[0].y; y < line[1].y+1; y++ {
			grid.bitmap[x][y] = true
		}
	}
}

func (grid grid) drawGround(ground coordinate) {
	y := ground.y
	for x := grid.bb.leftBottom.x; x < grid.bb.rightTop.x+1; x++ {
		grid.bitmap[x][y] = true
	}
}

func normalize(line Line) Line {
	if line[0].x < line[1].x || line[0].y < line[1].y {
		return line
	} else {
		return Line{line[1], line[0]}
	}
}

func printGrid(grid grid) {
	for j := grid.bb.leftBottom.y; j < grid.bb.rightTop.y+1; j++ {
		for i := grid.bb.leftBottom.x; i < grid.bb.rightTop.x+1; i++ {
			if grid.bitmap[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type grid struct {
	bitmap [][]bool
	bb     boundingBox
}

type path []coordinate
type Line [2]coordinate
type coordinate struct{ x, y int }
type boundingBox struct {
	leftBottom coordinate
	rightTop   coordinate
}

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

func makeGrid(bb boundingBox) grid {
	maxCoord := bb.rightTop

	maxX := maxCoord.x + 1
	maxY := maxCoord.y + 1

	bitmap := make([][]bool, maxX)
	for X := 0; X < maxX; X++ {
		bitmap[X] = make([]bool, maxY)
	}
	return grid{bitmap, bb}
}

type state int

const (
	falling     state = 1
	rest        state = 2
	endlessVoid state = 3
)

func simulateFalling(grid grid, startingCoord coordinate) int {
	var countRest int
	fallingSand := startingCoord
	var done bool
	for !done {
		state, nextCoord := next(grid, fallingSand)
		switch state {
		case falling:
			fallingSand = nextCoord
		case rest:
			grid.bitmap[fallingSand.x][fallingSand.y] = true
			countRest++
			fallingSand = startingCoord
			// restart the process
		case endlessVoid:
			done = true
		}
	}
	return countRest
}

func simulateFallingStep2(grid grid, startingCoord coordinate) int {
	var countRest int
	fallingSand := startingCoord
	var done bool
	for !done {
		state, nextCoord := next(grid, fallingSand)
		switch state {
		case falling:
			fallingSand = nextCoord
		case rest:
			if fallingSand == startingCoord {
				done = true
			} else {
				grid.bitmap[fallingSand.x][fallingSand.y] = true
				countRest++
				fallingSand = startingCoord
				// restart the process
			}
		case endlessVoid:
			// no op
			// panic("shouldn't happen")
		}
	}
	return countRest
}

func next(grid grid, c coordinate) (state, coordinate) {
	nextMoves := []coordinate{c.moveDown(), c.moveLeft(), c.moveRight()}
	for _, move := range nextMoves {
		switch {
		case grid.outOfBounds(move):
			return endlessVoid, c
		case grid.canMove(move):
			return falling, move
		}
	}
	return rest, c
}

func (g grid) canMove(c coordinate) bool {
	return !g.bitmap[c.x][c.y]
}

func (g grid) outOfBounds(c coordinate) bool {
	return !g.bb.within(c)
}

func getBoundingBox(paths []path) boundingBox {
	bb := paths[0].boundingBox()
	for _, p := range paths[1:] {
		bb = bb.union(p.boundingBox())
	}
	return bb
}

func (a boundingBox) union(b boundingBox) boundingBox {
	return boundingBox{a.leftBottom.min(b.leftBottom), a.rightTop.max(b.rightTop)}
}

func (a boundingBox) within(c coordinate) bool {
	return a.leftBottom.x <= c.x && c.x <= a.rightTop.x &&
		a.leftBottom.y <= c.y && c.y <= a.rightTop.y
}

func (p path) boundingBox() boundingBox {
	bb := p[0].boundingBox()
	for _, c := range p {
		bb = bb.union(c.boundingBox())
	}
	return bb
}

func (a coordinate) boundingBox() boundingBox {
	return boundingBox{a, a}
}

func (a coordinate) moveDown() coordinate {
	return coordinate{a.x, a.y + 1}
}

func (a coordinate) moveLeft() coordinate {
	return coordinate{a.x - 1, a.y + 1}
}

func (a coordinate) moveRight() coordinate {
	return coordinate{a.x + 1, a.y + 1}
}

func (a coordinate) min(b coordinate) coordinate {
	return coordinate{min(a.x, b.x), min(a.y, b.y)}
}

func (a coordinate) max(b coordinate) coordinate {
	return coordinate{max(a.x, b.x), max(a.y, b.y)}
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
