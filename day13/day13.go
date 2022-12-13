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
	case left[0] == '[' && right[0] == '[':
		l := splitList(left)
		r := splitList(right)
		return compareLists(l, r)
	// mixed types
	case left[0] == '[':
		r := makeList(right)
		return compare(left, r)
	default:
		l := makeList(left)
		return compare(l, right)
	}
}

func makeList(left string) string {
	return "[" + left + "]"
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

func compareLists(left, right []string) int {
	m := min(len(left), len(right))
	for i := 0; i < m; i++ {
		c := compare(left[i], right[i])

		// comparision is conclusive, return the result
		if c == -1 || c == 1 {
			return c
		}

		// items are the same, continue with the next item
	}
	switch {
	case len(left) < len(right):
		return -1
	case len(left) == len(right):
		return 0
	default:
		return 1
	}
}

func splitList(s string) []string {
	list := make([]string, 0)
	openBracesCount := 0
	v := ""

	for _, r := range s[1:] {
		switch {
		case r == '[':
			v += string(r)
			openBracesCount++
		case openBracesCount > 0 && r == ']':
			v += string(r)
			openBracesCount--
		case openBracesCount == 0 && r == ',' || r == ']':
			if v != "" {
				list = append(list, v)
				v = ""
			}
		default:
			v += string(r)
		}
	}
	return list
}

func isInt(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func sum(indices []int) int {
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

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
