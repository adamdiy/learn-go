package main

import (
	"fmt"
)

//onl two params
func sum(a, b float64) (res float64) {
	return a + b
}

func main() {
	fmt.Println(sum(1.1, 1.2))
}
