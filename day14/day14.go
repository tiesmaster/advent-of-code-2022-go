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
	bitmap := make([][]bool, maxCoord.x)
	for X := 0; X < maxCoord.x; X++ {
		bitmap[X] = make([]bool, maxCoord.y)
	}
	return grid{bitmap, bb}
}

func simulateFalling(grid grid, startingCoord coordinate) int {
	panic("unimplemented")
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

// func (a boundingBox) unionCoord(b coordinate) boundingBox {
// 	return a.union(b.boundingBox())
// }

// func (a coordinate) unionCoord(b coordinate) boundingBox {
// 	return boundingBox{a.min(b), a.max(b)}
// }

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
