package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

//takes a salutation object type Salutation
func Greet(salutation Salutation) {
	//calls CreateMessage function
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Yo", "Sup", "Ni Hao!")
	fmt.Println(message)
	fmt.Println(alternate)
}

//greeting is now a variadic function
func CreateMessage(name string, greeting ...string) (message string, alternate string) {

	fmt.Println("Len of name is:", len(name))
	fmt.Println("Len of greeting is:", len(greeting))

	fmt.Println(greeting[2])
	fmt.Println(name)

	message = greeting[2] + " " + name
	alternate = "HEY!" + " " + name
	return
}

func main() {
	var s = Salutation{"Don", "Hello"}
	//Here in main we pass in the s object
	Greet(s)
}
