package main

import (
	"fmt"
)

func main() {
	//when would you need to choose a float64
	var i float32

	fmt.Print("Please enter the numeric value you want to convert: ")
	_, err := fmt.Scan(&i)

	//How come when I enter a string nothing happens?
	if err != nil {
		fmt.Println("Error:", err)
	}
}
