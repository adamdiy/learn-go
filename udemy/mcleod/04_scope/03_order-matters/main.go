package main

import (
	"fmt"
)
var y = 42

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	//x needs to be moved above fmt as well
	//x := 42
}

//This won't work, needs to be moved to the top above func main
//var y = 42
