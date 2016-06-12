package main

import (
	"fmt"
)

func main() {
	a := 10 //int
	b := "golang" //string
	c := 4.17 //float
	d := true //bool

	//%T prints out the inferred type

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
