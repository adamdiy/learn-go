package main

import (
	"fmt"
)

//%v = format verb
func main() {
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 192.0 * 0.3783)
	fmt.Printf(" and I would be %v years old.\n", 32 * 365 / 687)
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 72)
	fmt.Printf("%-15v $%4v\n", "SpaceX", "94")
	fmt.Printf("%-14v $%4v\n", "SpaceX", "94")
	fmt.Printf("%-15v $%4v\n", "Virgin Galatic", "100")
}
