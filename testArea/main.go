package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2*z)
	}

	return z
}

func main() {
	// fmt.Println(Sqrt(2))
	math.Sqrt()

	for i := 0; i < 10; i++ {

		fmt.Printf("****************** CALCULATING SQRT FOR: %v *****************", i)
		fmt.Println()
		Sqrt(float64(i))

		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}
