package main

import (
	"fmt"
)

func main() {
	//Declare a slice called myCourses with a len of 5 and a capacity of 10
	//myCourses := make([]string, 5, 10)
	myCourses := []string{"Docker", "Puppet", "Python"}
	fmt.Printf("Length is: %d,\nCapacity is: %d", len(myCourses), cap(myCourses)) 
}
