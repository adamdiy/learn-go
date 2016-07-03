package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	//This allows the go routine to complete above, without it, this would just not return any output.
	var input string
	fmt.Scanln(&input)
}
