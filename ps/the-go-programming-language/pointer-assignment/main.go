package main

import (
	"fmt"
)

func main() {
	var message string = "Hello Go World"
	var greeting *string = &message

	//setting the actual pointer location for greeting to hi
	*greeting = "hi"
	fmt.Println(message, *greeting)
}
