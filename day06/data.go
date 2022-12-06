package day06

import (
	_ "embed"
)

var TestData = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

//go:embed data.txt
var Data string
