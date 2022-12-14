package main

import "fmt"

func main() {
	A, B := 2, 3
	toplevel := make([][]bool, A)
	for i := 0; i < A; i++ {
		toplevel[i] = make([]bool, B)
	}
	toplevel[1][2] = true
	fmt.Println(toplevel)
}
