package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

//2 - I don't even get why we need to do this? Why can't we just Println?
func Greet(salutation Salutation) {
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

//3
func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

//1
func main() {
	var s = Salutation{"Hello", "Bob"}
	s.name = "Bob"
	s.greeting = "Hello"
	Greet(s)
}
