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
			t.Errorf("CalculateSumOfPriorities(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCalculateSumOfPrioritiesStep2(t *testing.T) {
	cases := []struct {
		in string
		want int
	}{
			{TestData, 70},
			{Data, 2817},
	}
	for _, c := range cases {
		got := CalculateSumOfPrioritiesStep2(c.in)
		if got != c.want {
			t.Errorf("CalculateSumOfPrioritiesStep2(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}