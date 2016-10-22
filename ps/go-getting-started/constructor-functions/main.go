package main

import (
	"fmt"
)

//go doesn't have the concept of constructors so constructor functions are used instead.

func main() {
	//foo := &myStruct
	foo := newMyStruct()
	foo.myMap["bar"] = "bar"
	fmt.Println(foo)
}

type myStruct struct {
	myMap map[string]string
}

//constructor function starts here

func newMyStruct() *myStruct {
	result := myStruct{}
	result.myMap = map[string]string{}
	return &result
}
