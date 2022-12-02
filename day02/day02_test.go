package day02

import "testing"

func TestCalculateTotalScore(t *testing.T) {
	cases := []struct {
		in string
		want int
	}{
			{TestData, 15},
			{Data, 13924},
	}
	for _, c := range cases {
		got := CalculateTotalScore(c.in)
		if got != c.want {
			t.Errorf("CalculateTotalScore(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}