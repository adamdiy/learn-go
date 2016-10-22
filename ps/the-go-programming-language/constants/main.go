package main

import (
	"fmt"
)

type Salutation struct {
	name string
	greeting string
}

const (
	PI = 3.14
	Language = "Golang"
	//iota represents successive untyped integer constants
	//why it starts at 2? I don't know...
	//actually it makes sense since PI = 0, Language = 1
	A = iota
	B = iota
	C = iota
)

const (
	D = iota
	E
	F
)

func main() {
	var s = Salutation{}
	s.name = "Joe"
	s.greeting = "Howdy"

	fmt.Println(s.greeting, s.name)

	fmt.Println(PI, Language)
	fmt.Println(A,B,C)

	fmt.Println(D,E,F)
	
}
