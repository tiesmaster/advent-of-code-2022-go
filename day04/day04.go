package day04

import (
	"strconv"
	"strings"
)

type sectionAssignment struct {
	start int
	end   int
}

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

	return isOverlapping(first, second) || isOverlapping(second, first)
}

func arePairsOverlappingAsUnion(assignmentPair string) bool {
	first, second := parsePair(assignmentPair)

	return isOverlappingAsUnion(first, second) || isOverlappingAsUnion(second, first)
}

func parsePair(assignmentPair string) (sectionAssignment, sectionAssignment) {
	assignments := strings.Split(assignmentPair, ",")

	firstAssignment := parseAssignment(assignments[0])
	secondAssignment := parseAssignment(assignments[1])

	return firstAssignment, secondAssignment
}

func parseAssignment(assignment string) sectionAssignment {
	sections := strings.Split(assignment, "-")

	start, _ := strconv.Atoi(sections[0])
	end, _ := strconv.Atoi(sections[1])

	return sectionAssignment{start, end}
}

func isOverlapping(superSet, set sectionAssignment) bool {
	return superSet.start <= set.start && superSet.end >= set.end
}

func isOverlappingAsUnion(set1, set2 sectionAssignment) bool {
	return (set2.start <= set1.start && set1.start <= set2.end) ||
		(set2.start <= set1.end && set1.end <= set2.end)
}
