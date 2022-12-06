package day01

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestFindMostCalories(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 24000},
		{Data, 65912},
	}
	for _, c := range cases {
		got := FindMostCalories(c.in)
		if got != c.want {
			t.Errorf("FindMostCalories(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}

func TestFindTop3MostCalories(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 45000},
		{Data, 195625},
	}
	for _, c := range cases {
		got := FindTop3MostCalories(c.in)
		if got != c.want {
			t.Errorf("FindTop3MostCalories(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
