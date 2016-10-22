package main

import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module float64
)

func main() {
	fmt.Println("Name is set to", reflect.TypeOf(name))
}
