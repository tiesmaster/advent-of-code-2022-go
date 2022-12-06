package day06

import (
	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
	"testing"
)

func TestFindStartOfPacket(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 7},
		{Data, 7},
	}
	for _, c := range cases {
		got := FindStartOfPacket(c.in)
		if got != c.want {
			t.Errorf("FindStartOfPacket(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
