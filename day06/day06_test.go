package day06

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestFindFirstStartOfPacket(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 7},
		{Data, 1965},
	}
	for _, c := range cases {
		got := FindFirstStartOfPacket(c.in)
		if got != c.want {
			t.Errorf("FindFirstStartOfPacket(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}

func TestFindFirstStartOfMessage(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 19},
		{Data, 2773},
	}
	for _, c := range cases {
		got := FindFirstStartOfMessage(c.in)
		if got != c.want {
			t.Errorf("FindFirstStartOfMessage(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
