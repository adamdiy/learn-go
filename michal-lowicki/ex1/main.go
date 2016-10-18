package main

import (
	"fmt"
)

//takes an unlimited # of args
func sum(numbers ...float64) (res float64) {
	for _, number := range numbers {
		res += number
	}
	return
}

func main() {
	fmt.Println(sum(1.1, 2.2, 3.3, 4.4))
}
