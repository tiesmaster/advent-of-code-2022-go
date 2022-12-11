package day11

import (
	"sort"
	"strconv"
	"strings"
)

func Step01(notes string) int {
	const totalRounds = 20
	const reliefLevelLowersByThreeFold = true

	monkeys := parseMonkeys(notes)
	takeRounds(monkeys, totalRounds, reliefLevelLowersByThreeFold)
	return calculateMonkeyBusiness(monkeys)
}

func Step02(notes string) int {
	const totalRounds = 10_000
	const reliefLevelLowersByThreeFold = false

	monkeys := parseMonkeys(notes)
	takeRounds(monkeys, totalRounds, reliefLevelLowersByThreeFold)
	return calculateMonkeyBusiness(monkeys)
}

type monkey struct {
	items               []int
	worryLevelOperation func(int) int
	next                nextMonkeyDecider
	inspectionCount     int
}

type nextMonkeyDecider struct {
	testNumber  int
	trueMonkey  int
	falseMonkey int
}

func parseMonkeys(notes string) []monkey {
	monkeys := make([]monkey, 0)
	monkeyParts := strings.Split(notes, "\n\n")
	for _, block := range monkeyParts {
		monkeys = append(monkeys, parseMonkey(block))
	}
	return monkeys
}

func parseMonkey(block string) monkey {
	lines := strings.Split(block, "\n")
	items := parseItems(lines[1])
	worryLevelOperation := parseOperation(lines[2])
	next := parseTest(lines[3:])
	return monkey{
		items:               items,
		worryLevelOperation: worryLevelOperation,
		next:                next,
	}
}

func parseItems(s string) []int {
	items := make([]int, 0)
	parts := strings.Split(s[18:], ", ")
	for _, n := range parts {
		items = append(items, toInt(n))
	}
	return items
}

func parseOperation(s string) func(int) int {
	operationText := s[13:]
	switch {
	case operationText == "new = old * old":
		return func(old int) int {
			return old * old
		}
	case strings.Contains(operationText, "*"):
		operant := toInt(operationText[12:])
		return func(old int) int {
			return old * operant
		}
	case strings.Contains(operationText, "+"):
		operant := toInt(operationText[12:])
		return func(old int) int {
			return old + operant
		}
	}

	panic("cannot reach")
}

func parseTest(s []string) nextMonkeyDecider {
	testNumber := toInt(s[0][21:])
	trueMonkey := toInt(s[1][29:])
	falseMonkey := toInt(s[2][30:])
	return nextMonkeyDecider{testNumber, trueMonkey, falseMonkey}
}

func takeRounds(monkeys []monkey, totalRounds int, reliefLevelLowersByThreeFold bool) {
	for round := 0; round < totalRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := &monkeys[i]
			for _, worryLevel := range monkey.items {
				worryLevel := monkey.worryLevelOperation(worryLevel)
				if reliefLevelLowersByThreeFold {
					worryLevel = worryLevel / 3
				} else {
					if round > 20 {
						worryLevel = worryLevel - 4
					}
				}
				if worryLevel%monkey.next.testNumber == 0 {
					destinationMonkey := &monkeys[monkey.next.trueMonkey]
					destinationMonkey.items = append(destinationMonkey.items, worryLevel)
				} else {
					destinationMonkey := &monkeys[monkey.next.falseMonkey]
					destinationMonkey.items = append(destinationMonkey.items, worryLevel)
				}
				monkey.inspectionCount++
			}
			monkey.items = make([]int, 0)
		}
	}
}

func calculateMonkeyBusiness(monkeys []monkey) int {
	inspectionCounts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspectionCounts[i] = monkey.inspectionCount
	}
	sort.Slice(inspectionCounts, func(i, j int) bool {
		return inspectionCounts[i] > inspectionCounts[j]
	})

	return inspectionCounts[0] * inspectionCounts[1]
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
