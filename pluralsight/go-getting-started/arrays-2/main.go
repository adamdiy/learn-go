package main

import (
	"fmt"
)

func main() {
	slice := make([]float32, 100)
	slice[0] = 10.
	slice[1] = 11.
	slice[2] = 15.

	//slice short assignment
	mySlice := []float32{14., 15., 19.}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(slice)
}
