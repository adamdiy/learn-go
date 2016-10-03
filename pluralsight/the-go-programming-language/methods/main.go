package main

import "./greeting"

func main() {

	slice := []greeting.Salutation {
		{ "Bob", "Hello" },
		{ "Joe", "Hi!" },
		{ "Mary", "What is up?" },
	}

	slice = append(slice[:1], slice[2:]... )
	greeting.Greet(slice, greeting.CreatePrintFunction("!"), true, 150)
}
