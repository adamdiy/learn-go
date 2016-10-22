package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	fmt.Println(alternate)
	fmt.Println(message)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	//return greeting + " " + name, "HEY!" + name
	fmt.Println(len(greeting))
	message = greeting[0] + " " + name
	alternate = "HEY!" + name
	return
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
