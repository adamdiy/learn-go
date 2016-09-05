package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/tmp/test1.txt")

	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}
