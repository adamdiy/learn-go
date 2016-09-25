package main

import "fmt"

const (
	one_cm_mm  = 10.0
	one_inch_cm = 2.54
)

func main() {
	one_inch_in_mm := one_inch_cm / one_cm_mm * 100
	fmt.Println(one_inch_in_mm)
}
