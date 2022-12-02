package day02

import (
	_ "embed"
)

var TestData = `A Y
B X
C Z`

//go:embed data.txt
var Data string
