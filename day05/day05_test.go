package day05

import (
	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
	"testing"
)

func TestPerformRearrangementProcedure(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{TestData, "CMZ"},
		{Data, "JDTMRWCQJ"},
	}
	for _, c := range cases {
		got := PerformRearrangementProcedure(c.in)
		if got != c.want {
			t.Errorf("PerformRearrangementProcedure(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}
