package main

import "../greeting"

func main() {
	var s = greeting.Salutation{"Don", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true)
}
