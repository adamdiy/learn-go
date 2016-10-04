package main

import "./greeting"

func RenameToFrog(r greeting.Renameable) {
	r.Rename("Frog")
}

func main() {

	salutations := greeting.Salutations {
		{ "Bob", "Hello" },
		{ "Joe", "Hi!" },
		{ "Mary", "What is up?" },
	}
	salutations[0].Rename("John")
	//dereference = get the address of
	RenameToFrog(&salutations[0])

	salutations.Greet( greeting.CreatePrintFunction("!"), true, 6)
	//slice = append(slice[:1], slice[2:]... )
	//greeting.Greet(salutations, greeting.CreatePrintFunction("!"), true, 150)
}
