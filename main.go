package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func main() {
	fmt.Println("most calories [test data]: ", findMostCalories(testData))
	fmt.Println("most calories [data]: ", findMostCalories(data))

	fmt.Println("most calories (top 3) [test data]: ", findTop3MostCalories(testData))
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

	sort.Slice(totals, func (i, j int) bool {
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