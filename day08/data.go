package day08

import (
	_ "embed"
)

var TestData = `30373
25512
65332
33549
35390`

//go:embed data.txt
var Data string
