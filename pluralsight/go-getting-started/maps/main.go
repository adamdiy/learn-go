package main

import (
	"fmt"
)

func main() {
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[42] = "Foo"
	myMap[43] = "Bar"
	fmt.Println(myMap)

	//print a key that doesn't exist
	fmt.Println(myMap[99])
}
