package main

import "../greeting"

func main() {
	var s = greeting.Salutation{"012345678910", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!"), true)
	greeting.TypeSwitchTest(1.4)
}
