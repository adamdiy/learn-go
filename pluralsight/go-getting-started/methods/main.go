package main

func main() {
	mp := messagePrinter{"foo"}
	//println(mp.message)
	//println(mp.message)
	mp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}
