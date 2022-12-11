package main

import "fmt"

func main() {
	var hoi []int
	hoi = make([]int, 4)
	hoi[0] = 0
	hoi[1] = 1
	hoi[2] = 2
	hoi[3] = 3

	fmt.Println(hoi)
	doStuff(hoi)
	fmt.Println(hoi)
}

func doStuff(hoi []int) {
	hoi[0] = 3
}
