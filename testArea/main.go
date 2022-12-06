package main

import (
	"fmt"
)

func main() {
	b := [5]int{1, 2, 3, 4, 5}

	s := b[2:]
	printSlice(s)

	s = append(s, 99)
	printSlice(s)

	s[0] = 666

	printSlice(s)

	fmt.Println(b)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}