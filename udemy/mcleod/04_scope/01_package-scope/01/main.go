package main

import (
	"fmt"
)

// x is printable by func foo in this package
var x int = 42

func main() {
	fmt.Println(x)
	foo()
	// This is local to main, can't be printed below
	// block level scope
	//y := 17
}

func foo() {
	fmt.Println(x)
	//fmt.Println(y)
}
