package main

import (
	"fmt"
)

//type = object in other languages, here we are using a the string type but re-declaring it to Salutation
//type Salutation string

//If the type is capitalized then it's considered "public"

type Salutation struct {
	name string
	greeting string
}

func main() {
	var s = Salutation{"John", "Hello!"}
	//var message Salutation = "Hello World"
	//fmt.Println(message)
	fmt.Println(s.name)
	fmt.Println(s.greeting)
	fmt.Println(s.name, s.greeting)
	
}
