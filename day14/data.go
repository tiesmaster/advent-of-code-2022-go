package day14

import (
	_ "embed"
)

var TestData = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

//go:embed data.txt
var Data string
