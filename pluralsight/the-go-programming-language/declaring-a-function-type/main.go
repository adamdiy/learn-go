package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

//Function that takes a string and doesn't return anything
type Printer func(string) ()

//takes a salutation object type Salutation
//func Greet(salutation Salutation) {
//This is so weird, we are passing Println to the "do" func
//func Greet(salutation Salutation, do func(string)) {
func Greet(salutation Salutation, do Printer) {
	//calls CreateMessage function
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Yo", "Sup", "Ni Hao!")
	/*
	fmt.Println(message)
	fmt.Println(alternate)
	*/
	do(message)
	do(alternate)
}

//greeting is now a variadic function
func CreateMessage(name string, greeting ...string) (message string, alternate string) {

	fmt.Println("Len of name is:", len(name))
	fmt.Println("Len of greeting is:", len(greeting))

	fmt.Println(greeting[3])
	fmt.Println(name)

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

func main() {
	var s = Salutation{"Don", "Hello"}
	//Here in main we pass in the s object
	Greet(s, Println)
}
