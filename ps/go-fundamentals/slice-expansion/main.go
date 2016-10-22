package main

import (
	"fmt"
)

func main() {
	//create a slice of ints with a len of 1 and a capacity of 4.
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is %d Capacity is: %d", len(mySlice), cap(mySlice))
	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nLength is %d Capacity is: %d", len(mySlice), cap(mySlice))
	}
}
