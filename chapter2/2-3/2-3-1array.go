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
	for i := 1; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
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
