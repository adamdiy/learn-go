package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.0000000000
	b := 3

	fmt.Println("a is type:", reflect.TypeOf(a))
	fmt.Println("b is type:", reflect.TypeOf(b))

	//c := a + b
	c := int(a) + b

	fmt.Println("c is type:", reflect.TypeOf(c), "and has the value of:", c)

}
