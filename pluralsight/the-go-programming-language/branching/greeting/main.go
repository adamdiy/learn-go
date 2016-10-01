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
	switch {
		case name == "Bob": 
			prefix = "Mr "
		case name == "Joe", name == "Amy", len(name) == 10:
			if name == "Amy" {
				prefix = "Boss woman "
			} else {
				prefix = "Boss man "
			}
		case name == "Mary": 
			prefix = "Mrs "
		default: 
			prefix = "Sup Dude "
	}
	return
}

//empty interface
func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	//switch t := x.(type) {
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float")
		case string:
			fmt.Println("string")
		case Salutation:
			fmt.Println("salutation")
		default:
			fmt.Println("unknown")
		}
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
