package main

import "./greeting"

func main() {
	//var s = greeting.Salutation{"012345678910", "Hello"}

	var s []int
	s = make([]int, 3)
	s[0] = 1
	s[1] = 10
	s[2] = 500
	//should error out here.
	s[3] = 20


	slice := []greeting.Salutation {
		{ "Bob", "Hello" },
		{ "Joe", "Hi!" },
		{ "Mary", "What is up?" },
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("!"), true, 150)
	//greeting.TypeSwitchTest(1.4)
}
