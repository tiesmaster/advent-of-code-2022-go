package day13

import (
	"strconv"
	"strings"
	"unicode"
)

func Step01(distressSignal string) int {
	pairs := parseSignal(distressSignal)
	indices := indicesRightOrder(pairs)
	return sum(indices)
}

type pair [2]string

func parseSignal(distressSignal string) []pair {
	packetPairs := strings.Split(distressSignal, "\n\n")

	pairs := make([]pair, len(packetPairs))
	for i, pair := range packetPairs {
		lines := strings.Split(pair, "\n")
		pairs[i] = [2]string{lines[0], lines[1]}
	}
	return pairs
}

func indicesRightOrder(pairs []pair) []int {
	indices := make([]int, 0)
	for index, pair := range pairs {
		if compare(pair[0], pair[1]) == -1 {
			indices = append(indices, index)
		}
	}
	return indices
}


func compare(left, right string) int {
	switch {
	case isInt(left[0]) && isInt(right[0]):
		l, r := toInt(left), toInt(right)
		return compareInt(l, r)
		// case: both integer
		// x, y := 
		
	}

	panic("unimplemented")
}

func compareInt(left, right int) int {
	switch {
	case left == right:
		return 0
	case left < right:
		return -1
	default:
		return 1
	}
}

func isInt(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func sum(indices []int) int{
	sum := 0
	for _, index := range indices {
		sum += index
	}
	return sum
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}