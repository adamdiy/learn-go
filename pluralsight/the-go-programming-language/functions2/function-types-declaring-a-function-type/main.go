package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

//7
//func that takes a string but doesn't return anything
type Printer func(string) ()

//2 - I don't even get why we need to do this? Why can't we just Println?
//4 Greet now takes a function as a parameter
//8 replaced the nested func with Printer type above
func Greet(salutation Salutation, do Printer) {
	//We don't care about message so we just get rid of it
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	//fmt.Println(CreateMessage(salutation.name, salutation.greeting))
	//fmt.Println(message)
	//fmt.Println(alternate)
	do(message)
	do(alternate)
}

//3 //takes in name and greeting which are strings returns two strings
//greeting is a variadic function
//greeting now becomes a slice
func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	//return greeting + " " + name, "Hey!" + " " + name
	//fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "Hey!" + " " + name
	return
}

//5
func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

//1
//6
func main() {
	var s = Salutation{"Hello", "Bob"}
	s.name = "Bob"
	s.greeting = "Hello"
	//Greet now can takes a type + a func
	Greet(s, PrintLine)
}
