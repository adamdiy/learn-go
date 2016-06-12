package main

import (
	"fmt"
)

// x is printable by func foo in this package
var x int = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
