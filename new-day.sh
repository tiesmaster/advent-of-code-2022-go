#!/bin/bash

DAY_NUM=${1:-$(date +%d)}

mkdir day${DAY_NUM}
cd day${DAY_NUM}

cat <<EOF >day${DAY_NUM}.go
package day${DAY_NUM}

func Step01(data string) int {
	panic("unimplemented")
}
EOF

cat <<EOF >day${DAY_NUM}_test.go
package day${DAY_NUM}

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
package day${DAY_NUM}

import (
	_ "embed"
)

var TestData = \`\`

//go:embed data.txt
var Data string
EOF

cat <<EOF >data.txt
Hello day${DAY_NUM}
EOF
