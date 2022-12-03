package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// printTafelVan(10)

	// const s = "สวัสดี"
	const s = "A"
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	var i int
	i = int('A')
	fmt.Println(i)
}

func printTafelVan(tafel int) {
	fmt.Printf("Hier komt de tafel van %v", tafel)
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%2v x %v = %3v", i, tafel, i*tafel)
		fmt.Println()
	}
}
