package greeting

import "fmt"

type Salutation struct {
	Name string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
//	switch name {
	switch name {
		case "Bob": 
			prefix = "Mr "
		case "Joe": 
			prefix = "Bossman "
		case "Mary": 
			prefix = "Mrs "
		default: 
			prefix = "Dude"
	}
	return
	}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey!" + " " + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func PrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
