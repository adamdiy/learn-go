package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

func Greet(salutation Salutation) {
	/*
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
	*/
	//passing salutation.name and salutation.greeting to Createmessage, which simply returns
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}


func main() {
	var s = Salutation{}
	s.name = "Don"
	s.greeting = "Hello"
	Greet(s)
}
