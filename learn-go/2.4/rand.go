package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//pull a random number from 0-9 and add 1
	var num = rand.Intn(10) + 1
	fmt.Println(rand.Intn(100))
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
