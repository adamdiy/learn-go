package main

import (
	"fmt"
)

func main() {
	familyName := make(map[string]int)
	familyName["Ky"] = 32
	fmt.Println(familyName)

	otherFamilyNames := map[string]int {
		"Nguyen": 20,
		"Ngo": 30,
	}

	fmt.Printf("\nFamily names are: %v\nOther family names are: %v\n", familyName, otherFamilyNames)


}
