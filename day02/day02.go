package day02

import "strings"

func CalculateTotalScore(strategyGuide string) int {
	parts := strings.Split(strategyGuide, "\n")

	total := 0
	for _, s := range parts {
		total += calculateTotalForLine(s)
	}

	return total
}

func calculateTotalForLine(line string) int {
	opponent := parseShape(line[:1])
	me := parseShape(line[2:])
	outcome := calculateOutcome(me, opponent)

	pointsSelected := int(me)
	pointsOutcome := int(outcome)

	return pointsSelected + pointsOutcome
}

type Shape int

const (
	rock     Shape = 1
	paper    Shape = 2
	scissors Shape = 3
)

func parseShape(s string) Shape {
	if s == "A" || s == "X" {
		return rock
	}

	if s == "B" || s == "Y" {
		return paper
	}

	return scissors
}

type Outcome int

const (
	won     Outcome = 6
	draw    Outcome = 3
	loss Outcome = 0
)

func calculateOutcome(x Shape, y Shape) Outcome {
	if x == y {
		return draw
	}

	if x == rock && y == paper {
		return loss
	}

	if x == rock && y == scissors {
		return won
	}

	if x == paper && y == rock {
		return won
	}

	if x == paper && y == scissors {
		return loss
	}

	if x == scissors && y == rock {
		return loss
	}

	if x == scissors && y == paper {
		return won
	}

	panic("cannot reach")
}