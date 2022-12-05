package main

import (
	"fmt"
)

type stack []byte


func main() {
	sourcebytes := []byte{1, 2, 3, 8}
	destBytes := []byte{2}

	fmt.Println(sourcebytes, destBytes)

	sourcebytes, destBytes = move(sourcebytes, destBytes, 2)
	fmt.Println(sourcebytes, destBytes)

	sourcebytes, destBytes = move(sourcebytes, destBytes, 3)
	fmt.Println(sourcebytes, destBytes)
}

func move(sourceStack stack, destinationStack stack, moveCount int) (stack, stack) {
	totalCountToMove := min(moveCount, len(sourceStack))
	toMove := sourceStack[len(sourceStack) - totalCountToMove:]
	destinationStack = append(destinationStack, toMove...)
	return sourceStack[:len(sourceStack) - totalCountToMove], destinationStack
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}