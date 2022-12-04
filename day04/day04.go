package day04

import (
	"strconv"
	"strings"
)

func CalculateOverlappingAssignmentPairs(listOfAssignmentPairs string) int {
	assignmentPairs := strings.Split(listOfAssignmentPairs, "\n")

	countOverlaps := 0
	for _, pair := range assignmentPairs {
		if arePairsOverlapping(pair) {
			countOverlaps++
		}
	}

	return countOverlaps
}

func CalculateOverlappingAssignmentPairsAsUnions(listOfAssignmentPairs string) int {
	assignmentPairs := strings.Split(listOfAssignmentPairs, "\n")

	countOverlaps := 0
	for _, pair := range assignmentPairs {
		if arePairsOverlappingAsUnion(pair) {
			countOverlaps++
		}
	}

	return countOverlaps
}

func arePairsOverlapping(assignmentPair string) bool {
	first, second := parsePair(assignmentPair)

	return first&second == first || first&second == second
}

func arePairsOverlappingAsUnion(assignmentPair string) bool {
	first, second := parsePair(assignmentPair)

	return first&second > 0
}

func parsePair(assignmentPair string) (int, int) {
	assignments := strings.Split(assignmentPair, ",")

	firstAssignment := parseAssignment(assignments[0])
	secondAssignment := parseAssignment(assignments[1])

	return firstAssignment, secondAssignment
}

func parseAssignment(assignment string) int {
	sections := strings.Split(assignment, "-")

	start, _ := strconv.Atoi(sections[0])
	end, _ := strconv.Atoi(sections[1])

	totalBits := end - start + 1
	shiftLeft := start - 1

	i := 1<<totalBits - 1
	i = i << shiftLeft

	return i
}
