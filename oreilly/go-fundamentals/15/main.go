package main

import (
	"fmt"
)
//transmitter
func emit(c chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		c <- word
	}

	close(c)
}

func main() {
	wordChannel := make(chan string)

	//start a go routine
	//I guess this is like backgrounding a task
	go emit(wordChannel)

	//reciever
	/*
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	*/
	//explicitly receive
	word := <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word, ok := <-wordChannel
	fmt.Printf("%s\n", word, ok)
}
