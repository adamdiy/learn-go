package main

import (
	"fmt"
)

func main() {
	/**
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 57
	myArray[2] = 62
	**/
	myArray := [...]int{ 42, 57, 62}

	fmt.Println(myArray)
	fmt.Println(len(myArray))
}
