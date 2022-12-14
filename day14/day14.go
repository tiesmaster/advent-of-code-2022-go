package day14

import (
	"strconv"
	"strings"
)

func Step01(scan string) int {
	paths := parseScan(scan)
	startingCoord := coordinate{500, 0}

	gridBoundingBox := getBoundingBox(paths).union(startingCoord.boundingBox())
	grid := makeGrid(gridBoundingBox)
	return simulateFalling(grid, startingCoord)
}

type grid struct {
	bitmap [][]bool
	bb     boundingBox
}

type path []coordinate
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
