package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func FindMostCalories(data string) int {
	elvesInventory := strings.Split(data, "\n\n")
	maxCalories := 0
	for _, ev := range elvesInventory {
		maxCalories = max(maxCalories, calculateTotalCalories(ev))
	}

	return maxCalories
}

func calculateTotalCalories(elfInventory string) int {
	calories := strings.Split(elfInventory, "\n")
	sum := 0
	for _, c := range calories {
		cal, _ := strconv.Atoi(c)
		sum += cal
	}

	return sum
}

func FindTop3MostCalories(data string) int {
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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
