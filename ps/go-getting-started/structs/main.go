package main

import (
	"fmt"
)

func main() {
	//foo := myStruct{}
	//Composite literal format, cleaner way to create object.
	//foo := myStruct{"bar"}
	
	//create object on the heap, this doesn't allow initializationas above so field value must be set after.
	//foo := new(myStruct)
	//foo := myStruct{}
	//foo := new(myStruct)
	//foo.myField = "bar"

	foo := &myStruct{"bar"}

	//println(foo.myField)
	fmt.Println(foo)
}

type myStruct struct {
	myField string
}



