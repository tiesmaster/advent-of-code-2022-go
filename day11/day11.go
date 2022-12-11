package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const totalRounds = 20

func Step01(notes string) int {
	monkeys := parseMonkeys(notes)
	monkeys = takeRounds(monkeys, totalRounds)
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
	parts := strings.Split(s[16:], ", ")
	for _, n := range parts {
		items = append(items, toInt(n))
	}
	return items
}

func parseOperation(s string) func(int) int {
	operationText := s[11:]
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
	testNumber := toInt(s[0][19:])
	trueMonkey := toInt(s[1][27:])
	falseMonkey := toInt(s[2][28:])
	return nextMonkeyDecider{testNumber, trueMonkey, falseMonkey}
}

func takeRounds(monkeys []monkey, totalRounds int) []monkey {
	for round := 0; round < totalRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			// monkey := monkeys[i]
			for _, worryLevel := range monkeys[i].items {
				worryLevel := monkeys[i].worryLevelOperation(worryLevel)
				worryLevel = worryLevel / 3
				if worryLevel%monkeys[i].next.testNumber == 0 {
					destinationMonkey := &monkeys[monkeys[i].next.trueMonkey]
					destinationMonkey.items = append(destinationMonkey.items, worryLevel)
				} else {
					destinationMonkey := &monkeys[monkeys[i].next.falseMonkey]
					destinationMonkey.items = append(destinationMonkey.items, worryLevel)
				}
				monkeys[i].inspectionCount++
			}
			monkeys[i].items = make([]int, 0)
		}
		fmt.Println("Round: ", round)
		fmt.Println(monkeys)
}

	return monkeys
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
