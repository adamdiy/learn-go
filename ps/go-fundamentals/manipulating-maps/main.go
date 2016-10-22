package main

import (
	"fmt"	
)

func main() {
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	fmt.Println(testMap["A"])

	testMap["A"] = 100
	testMap["F"] = 1973
	fmt.Println(testMap["A"])
	fmt.Println(testMap["F"])
	delete(testMap, "F")
	fmt.Println(testMap["F"])
	
}
