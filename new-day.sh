#!/bin/bash

mkdir day09
cd day09

# touch day09{,_test}.go data.go data.txt

cat <<EOF >day09.go
package day09

func Step01(data string) int {
	panic("unimplemented")
}
EOF

cat <<EOF >day09_test.go
package day09

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestStep01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 666},
		// {Data, 1543},
	}
	for _, c := range cases {
		got := Step01(c.in)
		if got != c.want {
			t.Errorf("Step01(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
EOF

cat <<EOF >data.go
package day09

import (
	_ "embed"
)

var TestData = \`\`

//go:embed data.txt
var Data string
EOF

cat <<EOF >data.txt
Hello day09
EOF
