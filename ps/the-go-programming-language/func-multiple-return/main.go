package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

func Greet(salutation Salutation) {
	//message, alternate := CreateMessage(salutation.name, salutation.greeting)
	_, alternate := CreateMessage(salutation.name, salutation.greeting)
	//fmt.Println(CreateMessage(salutation.name, salutation.greeting))
	//fmt.Println(message)
	fmt.Println(alternate)
}

//return two strings
func CreateMessage(name, greeting string) (string, string) {
	//return two returns
	return greeting + " " + name, "HEY" + name
}

func main() {
	var s = Salutation{"Don", "Hello"}
	Greet(s)
	
}
