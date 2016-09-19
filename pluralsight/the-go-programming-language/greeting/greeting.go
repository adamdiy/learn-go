package greeting 

import (
	"fmt"
)

type Salutation struct {
	Name string
	Greeting string
}

//type Printer that takes in a func
type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting, "Yo", "Sup", "Ni Hao!")
	if isFormal {
		do(message)
	}
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message = greeting[0] + " " + name
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

