package main

import (
	"fmt"
)

func sort(array []float32) {
	var temp float32
	for i := 0; i < len(array); i++ {
		if array[i] > array[i+1] {
			array[i] = temp
			array[i+1] = array[i]
			temp = array[i+1]
		}
	}
	for j := 0; j < len(array); j++ {
		fmt.Println(array[j])
	}
}
func main() {

	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	sort(balance)

}
