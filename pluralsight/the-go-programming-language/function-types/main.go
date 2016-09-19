package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

//Notice the Greet func accepts a do func with a string as a input
func Greet(salutation Salutation, do func(string)) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	/*
	fmt.Println(message)
	fmt.Println(alternate)
	*/
	do(message)
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message = greeting[1] + " " + name
	alternate = "HEY!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func Printme(someString string) {
	fmt.Println(someString)
}

func main() {
	var s = Salutation{"Don", "Hello!"}
	//passing a print fuction to the Greet function
	Greet(s, Print)
	//Greet(s, Printme("yabadaba"))
	Printme("Yayaaaa")
	
}
