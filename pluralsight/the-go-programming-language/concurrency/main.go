package main

import (
	"./greeting"
	_"time"
)

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

	done := make(chan bool)

	//Kick off independent thread with this
	go func() {
		salutations.Greet( greeting.CreatePrintFunction("<C>"), true, 6)
		done <- true
	}()
	salutations.Greet( greeting.CreatePrintFunction("!"), true, 6)
	//time.Sleep(100 * time.Millisecond)
	<- done
	//slice = append(slice[:1], slice[2:]... )
	//greeting.Greet(salutations, greeting.CreatePrintFunction("!"), true, 150)
}
