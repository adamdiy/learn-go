package main

import (
	"fmt"
)

//multiple var inside a block
var (
	name, course string
	module float64
)

func main() {
	fmt.Println("Name is set to", name)
	fmt.Println("Module is set to", module)
}
