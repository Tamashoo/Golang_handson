package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a price")
	price, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	p := float64(price)
	fmt.Println(int(p * 1.1))
}
