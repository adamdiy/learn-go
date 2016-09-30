package main

import "../greeting"


func main() {
	var s = greeting.Salutation{"Hello", "Bob"}
	greeting.Greet(s, greeting.CreatePrintFunction("!"), false)
}
