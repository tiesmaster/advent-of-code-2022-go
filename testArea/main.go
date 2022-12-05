package main

import (
	"fmt"
)

func main() {
	sourcebytes := []byte{1, 2, 3, 8}
	destBytes := []byte{2}

	fmt.Println(sourcebytes, destBytes)

	sourcebytes, destBytes = move(sourcebytes, destBytes, 2)
	fmt.Println(sourcebytes, destBytes)
}

func move(sourcebytes []byte, destBytes []byte, q int) ([]byte, []byte) {
	toMove := sourcebytes[len(sourcebytes)-q:]
	destBytes = append(destBytes, toMove...)
	return sourcebytes[:q], destBytes
}
