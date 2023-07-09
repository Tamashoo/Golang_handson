package main

import (
	"fmt"
	"hello"
	"strconv"
	"strings"
)

func main() {
	t := 0
	x := hello.Input("input data")
	ar := strings.Split(x, " ")
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			goto error
		}
		t += n
	}
	fmt.Println("total:", t)
	return

error:
	fmt.Println("ERROR!")
}
