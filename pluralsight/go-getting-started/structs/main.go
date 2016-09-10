package main

import "fmt"

func main() {
	//foo := myStruct{}
	//foo.myField = "bar"

	//foo := myStruct{"bar"}
	//foo := new(myStruct)
	//foo := myStruct{}
	//foo.myField = "bar"

	foo := &myStruct{"bar"}

	fmt.Println(foo)
	println(foo.myField)
}

type myStruct struct {
	myField string
}
