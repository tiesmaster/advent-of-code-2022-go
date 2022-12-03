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
	shapeParsingMapping = map[string]Shape{
		"A": rock,
		"B": paper,
		"C": scissors,

		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	outcomeParsingMapping = map[string]Outcome{
		"X": loss,
		"Y": draw,
		"Z": won,
	}
	outcomeMapping = map[[2]Shape]Outcome{
		{rock, rock}:         draw,
		{rock, paper}:        loss,
		{rock, scissors}:     won,
		{paper, rock}:        won,
		{paper, paper}:       draw,
		{paper, scissors}:    loss,
		{scissors, rock}:     loss,
		{scissors, paper}:    won,
		{scissors, scissors}: draw,
	}
)

func parseShape(s string) Shape {
	return shapeParsingMapping[s]
}

func parseOutcome(s string) Outcome {
	return outcomeParsingMapping[s]
}

func calculateOutcome(x Shape, y Shape) Outcome {
	return outcomeMapping[[2]Shape{x, y}]
}

func calculateShapeToChoose(opponentShape Shape, desiredOutcome Outcome) Shape {
	switch desiredOutcome {
	case won:
		switch opponentShape {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		}
	case draw:
		return opponentShape
	case loss:
		switch opponentShape {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	}

	panic("cannot reach")
}
