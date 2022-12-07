package day07

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestStep01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 95437},
		{Data, 1307902},
	}
	for _, c := range cases {
		got := Step01(c.in)
		if got != c.want {
			t.Errorf("Step01(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}


func TestStep02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 24933642},
		// {Data, 1307902},
	}
	for _, c := range cases {
		got := Step02(c.in)
		if got != c.want {
			t.Errorf("Step02(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}