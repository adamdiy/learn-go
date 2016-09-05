package main

import (
	"fmt"
)

//if must be boolean expression
func main() {
	firstRank := "39" //Docker Deep dive
	secondRank := "614" //Docker clustering

	if firstRank < secondRank {
		fmt.Println("\nThe First course is doing better" +
		" than second course")
	} else if firstRank > secondRank {
		fmt.Println("\nOh deer... your first" +
		"course must be doing abysmally!")
	} else {
		fmt.Println("\nBoth courses are either" + 
		"the same or something weird is going on")
	}

}
