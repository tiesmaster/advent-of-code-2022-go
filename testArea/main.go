package main

import (
	"fmt"
)

type stack []byte


func main() {
	bytes := stack{6, 8, 3}
	fmt.Println("original bytes: ", bytes)

	bytes, b := pop(bytes)

	fmt.Println("byte: ", b)
	fmt.Println("bytes: ", bytes)

	bytes = push(bytes, 123)
	fmt.Println("bytes: ", bytes)
}

func push(bytes stack, item byte) stack {
	return append(bytes, item)
}

func pop(bytes stack) (stack, byte) {
	n := len(bytes)
	ret := bytes[n-1]

	bytes = bytes[:n-1]

	return bytes, ret
}
