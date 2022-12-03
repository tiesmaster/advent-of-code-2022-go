package main

import "fmt"

func main() {
	printTafelVan(10)
}

func printTafelVan(tafel int) {
	fmt.Printf("Hier komt de tafel van %v", tafel)
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%2v x %v = %3v", i, tafel, i * tafel)
		fmt.Println()
	}
}