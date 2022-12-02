package main

import "fmt"

func main() {

	line := "Yolo"
	fmt.Println(line[:1])

	var yolo Shape;
	yolo = rock;

	var i int

	i = int(yolo)

	fmt.Println(i)
}

type Shape int

const (
	rock     Shape = 1
	paper    Shape = 2
	scissors Shape = 3
)