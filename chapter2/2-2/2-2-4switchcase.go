package main

import (
	"fmt"
)

func main() {
	x := 5
	switch x {
	case f(1):
		fmt.Println("First case")
	case f(2):
		fmt.Println("Second case")
	case f(3):
		fmt.Println("Third case")
	default:
		fmt.Println("Default case")
	}
}

func f(n int) int {
	fmt.Println("No,", n)
	return n
}
