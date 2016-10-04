package main

import "./greeting"

func main() {

	salutations := greeting.Salutations {
		{ "Bob", "Hello" },
		{ "Joe", "Hi!" },
		{ "Mary", "What is up?" },
	}
	salutations[0].Rename("John")

	salutations.Greet( greeting.CreatePrintFunction("!"), true, 6)
	//slice = append(slice[:1], slice[2:]... )
	//greeting.Greet(salutations, greeting.CreatePrintFunction("!"), true, 150)
}
