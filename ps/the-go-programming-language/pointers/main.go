package main

import (
	"fmt"
)

func main() {
	message := "Hello Go World!"

	// set greeting to the memory location address of message
	var greeting *string = &message

	fmt.Println(message, greeting)
	//dereference memory location, should print back actual string message
	fmt.Println(*greeting)
}
