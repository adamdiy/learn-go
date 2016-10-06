package main

import (
	"./greeting"
	_"time"
	"fmt"
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

	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	c := make(chan greeting.Salutation)
	// Call a goroutine that will fill the channel
	for s := range c {
		fmt.Println(s.Name)
	}
}
