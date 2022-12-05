package main

import (
	"fmt"
)

func main() {
	bytes := []byte{6, 8, 3}
	fmt.Println("original bytes: ", bytes)

	bytes, b := pop(bytes)

	fmt.Println("byte: ", b)
	fmt.Println("bytes: ", bytes)

	bytes = push(bytes, 123)
	fmt.Println("bytes: ", bytes)
}

func pop(bytes []byte) ([]byte, byte) {
	n := len(bytes)
	ret := bytes[n-1]

	bytes = bytes[:n-1]

	return bytes, ret
}

func push(bytes []byte, item byte) []byte {
	return append(bytes, item)
}
