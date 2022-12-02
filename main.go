package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tiesmaster/advent-of-code-2022-go/day02"
)

func main() {
	fmt.Println("day 01: most calories [test data]: ", findMostCalories(testData))
	fmt.Println("day 01: most calories [data]: ", findMostCalories(data))

	fmt.Println("day 01: most calories (top 3) [test data]: ", findTop3MostCalories(testData))
	fmt.Println("day 01: most calories (top 3) [data]: ", findTop3MostCalories(data))

	fmt.Println("day 02: total score [test data]", day02.CalculateTotalScore(day02.TestData))
	fmt.Println("day 02: total score [data]", day02.CalculateTotalScore(day02.Data))

	fmt.Println("day 02: total score (step 2) [test data]", day02.CalculateTotalScoreStep2(day02.TestData))
	fmt.Println("day 02: total score (step 2) [data]", day02.CalculateTotalScoreStep2(day02.Data))
}

func findTop3MostCalories(data string) int {
	parts := strings.Split(data, "\n")

	totals := make([]int, 0)
	sum := 0

	for _, s := range parts {
		if s == "" {
			totals = append(totals, sum)
			sum = 0
		} else {
			i, _ := strconv.Atoi(s)
			sum += i
		}
	}

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	top3 := totals[:3]

	return sumTotals(top3)
}

func sumTotals(totals []int) int {
	sum := 0
	for _, s := range totals {
		sum += s
	}

	return sum
}

func printTotals(totals []int) {
	for _, s := range totals {
		fmt.Println(s)
	}
}

func findMostCalories(data string) int {
	parts := strings.Split(data, "\n")

	max := 0
	sum := 0

	for _, s := range parts {
		if s == "" {
			max = Max(max, sum)
			sum = 0
		} else {
			i, _ := strconv.Atoi(s)
			sum += i
		}
	}

	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
