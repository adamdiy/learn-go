package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

//2 - I don't even get why we need to do this? Why can't we just Println?
func Greet(salutation Salutation) {
	//fmt.Println(salutation.name)
	//fmt.Println(salutation.alternate)
	//We don't care about message so we just get rid of it
	//message, alternate := CreateMessage(salutation.name, salutation.greeting)
	//fmt.Println(CreateMessage(salutation.name, salutation.greeting))
	//fmt.Println(message)
	//fmt.Println(alternate)
}

//3 //takes in name and greeting which are strings returns two strings
func CreateMessage(name, greeting string) (message string, alternate string) {
	//return greeting + " " + name, "Hey!" + name
	//store the return values into variables
	message = greeting + " " + name
	alternate = "HEY!" + name
	return
}

//1
func main() {
	var s = Salutation{"Bob", "Hello"}
	s.name = "Bob"
	s.greeting = "Hello"
	Greet(s)
}
