package main

func main() {
	//message := "Hello"
	sayHello("message", "message2")
}

func sayHello(messages ...string) {
	for _, msg := range messages {
		println(msg)
	}
}
