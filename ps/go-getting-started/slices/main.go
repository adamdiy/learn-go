package main

import (
	"fmt"
)

func main() {
	//myArray is an array of ints with 3 values
	myArray := [...]int{ 42, 57, 62}
	//create a slice of all values in myArray : means to assign all values of the array into the slice.
	mySlice := myArray[:]
	//ignore the 0 element and give me everything after
	mySlice = myArray[1:]
	//give me everything from the 0 element to the ending element of 3
	mySlice = myArray[:3]

	//append an additional value to mySlice
	mySlice = append(mySlice, 100)

	fmt.Println(myArray)
	fmt.Println(mySlice)
}
