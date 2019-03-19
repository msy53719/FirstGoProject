package main

import (
	"fmt"
)

//const (
//	a = iota
//	b
//	c
//	d = "ha"
//	e
//	f = 100
//	h
//	i = iota
//	g
//)
func swap(x int, y int) int {
	return x + y
}
func main() {
	//var str string
	//str := "asfafa"
	//const b string = "sdfsf"
	//fmt.Printf("Hello  GO !", b)
	//fmt.Println(a, b, c, d, e, f, h, i, g)
	o, p := 1, 2
	if o == p {
		fmt.Println("执行aaaa")
	} else {
		fmt.Println("执行BB")
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	x, y := 1, 5
	fmt.Println(swap(x, y))

}
