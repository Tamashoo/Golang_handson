package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	t := 0
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		goto error
	}

	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
	return

error:
	fmt.Println("ERROR!")
}
