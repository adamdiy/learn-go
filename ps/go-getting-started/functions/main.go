package main

func main() {
	message := "Hello"
	sayHello(&message)
	println(message)
}

func sayHello(msg *string) {
	//println("Hello")
	//fmt.Println(msg)
	//print out memeory location
	//println(msg)
	//dereference the pointer
	println(*msg)
	//re-assign "Hello Go" to the msg
	*msg = "Hello Go"
}
