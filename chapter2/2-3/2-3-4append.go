package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30}
	b := a[0:2]
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
}
