package day03

import (
	"strings"
)

func CalculateSumOfPriorities(listOfContents string) int {

	rucksacks := strings.Split(listOfContents, "\n")

	sum := 0
	for _, rucksack := range rucksacks {
		sum += calculatePriority(rucksack)
	}

	return sum
}

func calculatePriority(rucksack string) int {
	sizeOfCompartment := len(rucksack) / 2

	compartment1 := rucksack[:sizeOfCompartment]
	compartment2 := rucksack[sizeOfCompartment:]

	for _, item := range compartment1 {
		if strings.Contains(compartment2, string(item)) {
			return calculatePriorityOfItem(item)
		}
	}

	panic("cannot reach")
}

func calculatePriorityOfItem(item rune) int {
	i := int(item)
	if i <= 'Z' {
		return i - 'A' + 27
	} else {
		return i - 'a' + 1
	}
}