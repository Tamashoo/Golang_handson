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
	fmt.Print(x + "は、")
	if price%2 == 0 {
		fmt.Println("偶数です。")
	} else {
		fmt.Println("奇数です。")
	}
}
