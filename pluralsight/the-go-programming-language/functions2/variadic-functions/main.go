package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

//2 - I don't even get why we need to do this? Why can't we just Println?
func Greet(salutation Salutation) {
	//We don't care about message so we just get rid of it
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	//fmt.Println(CreateMessage(salutation.name, salutation.greeting))
	fmt.Println(message)
	fmt.Println(alternate)
}

//3 //takes in name and greeting which are strings returns two strings
//greeting is a variadic function
//greeting now becomes a slice
func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	//return greeting + " " + name, "Hey!" + " " + name
	message = greeting[0] + " " + name
	alternate = "Hey!" + " " + name
	return
}

//1
func main() {
	var s = Salutation{"Hello", "Bob"}
	s.name = "Bob"
	s.greeting = "Hello"
	Greet(s)
}
