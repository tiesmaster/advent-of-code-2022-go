package day04

import (
	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
	"testing"
)

func TestCalculateOverlappingAssignmentPairs(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 2},
		{Data, 657},
	}
	for _, c := range cases {
		got := CalculateOverlappingAssignmentPairs(c.in)
		if got != c.want {
			t.Errorf("CalculateOverlappingAssignmentPairs(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
