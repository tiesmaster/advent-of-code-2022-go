package day11

import (
	"strconv"
	"strings"
)

const totalRounds = 20

func Step01(notes string) int {
	monkeys := parseMonkeys(notes)
	takeRounds(&monkeys, totalRounds)
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
	testNumber := toInt(s[0][20:])
	trueMonkey := toInt(s[1][27:])
	falseMonkey := toInt(s[2][28:])
	return nextMonkeyDecider{testNumber, trueMonkey, falseMonkey}
}

func takeRounds(monkeys *[]monkey, i int) {
	panic("unimplemented")
}

func calculateMonkeyBusiness(monkeys []monkey) int {
	panic("unimplemented")
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
