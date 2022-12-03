package day02

import (
	"strings"
)

func CalculateTotalScore(strategyGuide string) int {
	parts := strings.Split(strategyGuide, "\n")

	total := 0
	for _, s := range parts {
		total += calculateTotalForLine(s)
	}

	return total
}

func CalculateTotalScoreStep2(strategyGuide string) int {
	parts := strings.Split(strategyGuide, "\n")

	total := 0
	for _, s := range parts {
		total += calculateTotalStep2ForLine(s)
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

func calculateTotalStep2ForLine(line string) int {
	opponent := parseShape(line[:1])
	desiredOutcome := parseOutcome(line[2:])

	me := calculateShapeToChoose(opponent, desiredOutcome)

	pointsSelected := int(me)
	pointsOutcome := int(desiredOutcome)

	return pointsSelected + pointsOutcome
}

type Shape int
type Outcome int

const (
	rock     Shape = 1
	paper    Shape = 2
	scissors Shape = 3
)

const (
	won  Outcome = 6
	draw Outcome = 3
	loss Outcome = 0
)

var (
	shapeMapping = map[string]Shape{
		"A": rock,
		"B": paper,
		"C": scissors,

		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	outcomeMapping = map[string]Outcome{
		"X": loss,
		"Y": draw,
		"Z": won,
	}
)

func parseShape(s string) Shape {
	return shapeMapping[s]
}

func parseOutcome(s string) Outcome {
	return outcomeMapping[s]
}

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

func calculateShapeToChoose(opponentShape Shape, desiredOutcome Outcome) Shape {
	if desiredOutcome == draw {
		return opponentShape
	}

	if desiredOutcome == won {
		if opponentShape == rock {
			return paper
		}

		if opponentShape == paper {
			return scissors
		}

		return rock
	}

	if desiredOutcome == loss {
		if opponentShape == rock {
			return scissors
		}

		if opponentShape == scissors {
			return paper
		}

		return rock
	}

	panic("cannot reach")
}
