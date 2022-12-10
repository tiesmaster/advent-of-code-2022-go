package day10

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestStep01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 13140},
		{Data, 17020},
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
		want string
	}{
		{TestData, TestDataWant},
		// {Data, DataWant},
	}
	for _, c := range cases {
		got := Step02(c.in)
		if got != c.want {
			t.Errorf("Step02(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
