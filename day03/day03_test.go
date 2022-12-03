package day03

import "testing"

func TestCalculateSumOfPriorities(t *testing.T) {
	cases := []struct {
		in string
		want int
	}{
			{TestData, 157},
			{Data, 8185},
	}
	for _, c := range cases {
		got := CalculateSumOfPriorities(c.in)
		if got != c.want {
			t.Errorf("CalculateTotalScore(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}