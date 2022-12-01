package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("most calories (test data): ", findMostCalories(testData))
	fmt.Println("most calories (data): ", findMostCalories(data))
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