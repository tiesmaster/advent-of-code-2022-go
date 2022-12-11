package day11

import (
	"testing"

	testhelpers "github.com/tiesmaster/advent-of-code-2022-go/testHelpers"
)

func TestStep01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 10605},
		{Data, 57348},
	}
	for _, c := range cases {
		got := Step01(c.in)
		if got != c.want {
			t.Errorf("Step01(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}

func _TestStep02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{TestData, 2713310158},
		// {Data, 57348},
	}
	for _, c := range cases {
		got := Step02(c.in)
		if got != c.want {
			t.Errorf("Step02(%v) == %v, want %v", testhelpers.TrimInput(c.in), got, c.want)
		}
	}
}

func _TestStep02Inspections(t *testing.T) {
	cases := []struct {
		in   int
		want [4]int
	}{
		{1, [4]int{2, 4, 3, 6}},
		{20, [4]int{99, 97, 8, 103}},
		{1000, [4]int{5204, 4792, 199, 5192}},
		{2000, [4]int{10419, 9577, 392, 10391}},
		{3000, [4]int{15638, 14358, 587, 15593}},
		{4000, [4]int{20858, 19138, 780, 20797}},
		{5000, [4]int{26075, 23921, 974, 26000}},
		{6000, [4]int{31294, 28702, 1165, 31204}},
		{7000, [4]int{36508, 33488, 1360, 36400}},
		{8000, [4]int{41728, 38268, 1553, 41606}},
		{9000, [4]int{46945, 43051, 1746, 46807}},
		{10000, [4]int{52166, 47830, 1938, 52013}},
	}
	for _, c := range cases {
		got := Step02Inspections(c.in)
		if got != c.want {
			t.Errorf("Step02Inspections(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func Step02Inspections(rounds int) [4]int {
	const reliefLevelLowersByThreeFold = false
	monkeys := parseMonkeys(TestData)
	monkeys = takeRounds(monkeys, rounds, reliefLevelLowersByThreeFold)

	var inspections [4]int
	for i := 0; i < 4; i++ {
		inspections[i] = monkeys[i].inspectionCount
	}
	return inspections
}
