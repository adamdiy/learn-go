package main

import (
	"fmt"
)

func f(names ...string) {
	fmt.Printf("value: %#v\n", names)
	fmt.Printf("length: %#v\n", len(names))
	fmt.Printf("capacity: %#v\n", cap(names))
	fmt.Printf("type: %T\n", names)
}

func main() {
	f("one", "two", "three")
}
