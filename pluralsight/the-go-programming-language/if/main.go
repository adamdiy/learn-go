package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

//type Printer that takes in a func
type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Yo", "Sup", "Ni Hao!")
	do(message)
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message = greeting[3] + " " + name
	alternate = "HEY!" + " " + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func Println(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {
	var s = Salutation{"Don", "Hello"}
	Greet(s, CreatePrintFunction("!!!"))
}
