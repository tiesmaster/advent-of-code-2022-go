package main

import "fmt"

func main() {

	printTafelVan(4)
}

func printTafelVan(tafel int) {
	fmt.Println("Hier komt de tafel van", tafel)
	fmt.Println()
	for i := tafel; i <= 10*tafel; i += tafel {
		fmt.Println("1 x 4 =", i)
	}

}