package day12

import (
	_ "embed"
)

var TestData = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

//go:embed data.txt
var Data string
