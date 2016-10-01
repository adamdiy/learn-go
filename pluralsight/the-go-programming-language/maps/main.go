package main

import "./greeting"

func main() {
	//var s = greeting.Salutation{"012345678910", "Hello"}

	slice := []greeting.Salutation {
		{ "Bob", "Hello" },
		{ "Joe", "Hi!" },
		{ "Mary", "What is up?" },
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("!"), true, 150)
	//greeting.TypeSwitchTest(1.4)
}
