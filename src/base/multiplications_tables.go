package main

import (
	"fmt"
)

func test() {
	var j, i int
	for i = 1; i < 10; i++ {
		for j = 1; j <= i; j++ {
			//fmt.Printf("%d \n", i, "*", j, "=", i*j)
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
func main() {
	test()
}
