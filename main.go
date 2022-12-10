package main

import "fmt"

func main() {
	var buffer videoBuffer
	
	// buffer[0][1] = 123
	row, col := 1, 2
	drawPixel(&buffer, row, col)

	fmt.Println(buffer)
}

const (
	numRows    = 6
	numColumns = 40
)

type videoBuffer [numRows][numColumns]int

func drawPixel(buffer *videoBuffer, row, col int) {
	buffer[row][col] = 123
}
