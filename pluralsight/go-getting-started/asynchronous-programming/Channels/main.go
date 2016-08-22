package main

import "time"
import "runtime"

func main() {
	runtime.GOMAXPROCS(8)

	go abcGen()
	time.Sleep(100 * time.Millisecond)
	println("This comes first!")
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
