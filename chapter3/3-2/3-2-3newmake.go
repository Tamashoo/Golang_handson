package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 5, 5)
	fmt.Println(taro)
}

func rev(md *Mydata) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
