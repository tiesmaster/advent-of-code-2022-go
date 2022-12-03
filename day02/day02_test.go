package day02

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestCalculateTotalScore(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 15},
		{Data, 13924},
	}
	for _, c := range cases {
		got := CalculateTotalScore(c.in)
		if got != c.want {
			t.Errorf("CalculateTotalScore(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}

func TestCalculateTotalScoreStep2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 12},
		{Data, 13448},
	}
	for _, c := range cases {
		got := CalculateTotalScoreStep2(c.in)
		if got != c.want {
			t.Errorf("CalculateTotalScoreStep2(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
