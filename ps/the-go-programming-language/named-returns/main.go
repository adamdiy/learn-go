package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(alternate)
	fmt.Println(message)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	//return greeting + " " + name, "HEY!" + name
	message = greeting + " " + name
	alternate = "HEY!" + name
	return
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
