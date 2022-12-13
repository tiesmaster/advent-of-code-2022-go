package day13

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestStep01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 13},
		// {Data, 1543},
	}
	for _, c := range cases {
		got := Step01(c.in)
		if got != c.want {
			t.Errorf("Step01(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}

func TestCompare(t *testing.T) {
	cases := []struct {
		in   [2]string
		want int
	}{
		{[2]string{"1", "2"}, -1},
		{[2]string{"1", "1"}, 0},
		{[2]string{"2", "1"}, 1},
		{[2]string{"123", "456"}, -1},
		{[2]string{"[1]", "[2]"}, -1},
		{[2]string{"[2]", "[1]"}, 1},
		{[2]string{"[1]", "[1]"}, 0},
		{[2]string{"[1,2,3]", "[2,2,2]"}, -1},
		{[2]string{"[1,2,3]", "[2]"}, -1},
		{[2]string{"[1,2,3]", "[1]"}, 1},
	}
	for _, c := range cases {
		got := compare(c.in[0], c.in[1])
		if got != c.want {
			t.Errorf("compare(%v, %v) == %v, want %v", c.in[0], c.in[1], got, c.want)
		}
	}
}