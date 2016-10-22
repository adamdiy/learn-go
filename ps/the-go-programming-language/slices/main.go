package main

import "./greeting"

func main() {
	//var s = greeting.Salutation{"012345678910", "Hello"}

	/*
	var s []int
	s = make([]int, 3)
	s[0] = 1
	s[1] = 10
	s[2] = 500
	//should error out here.
	//s[3] = 20
	*/

	//short hand assignment
	//s := []int { "1, 10, 500, 25 }

	slice := []greeting.Salutation {
		{ "Bob", "Hello" },
		{ "Joe", "Hi!" },
		{ "Mary", "What is up?" },
	}

	//we want to extract joe so we specify the start index + end index
	//slice = slice[1:2]
	//slice = slice[:2]
	//slice = append(slice, greeting.Salutation{"Frank", "Heyo "})
	//expand the slice with ...
	//slice = append(slice, slice...)

	//deleting
	slice = append(slice[:1], slice[2:]... )


	greeting.Greet(slice, greeting.CreatePrintFunction("!"), true, 150)

	//greeting.TypeSwitchTest(1.4)
}
