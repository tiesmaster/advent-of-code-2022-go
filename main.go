package main

import "fmt"

func main() {
	l := []string{"1", "2", "3"}
	s := fmt.Sprint(l)
	fmt.Println("[" + s + "]")

	s1 := ""
	fmt.Println(len(s1))
}
