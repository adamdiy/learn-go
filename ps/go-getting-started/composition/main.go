package main

func main() {
	//mp := messagePrinter{"foo"}
	//mp := enhancedMessagePrinter{"foo"}
	emp := enhancedMessagePrinter{}
	emp.message = "foo"
	emp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
