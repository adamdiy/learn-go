package main

import ( 
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	go abcGen()	

	println("this comes first")

	time.Sleep(100 * time.Millisecond)	
}

func abcGen() {
	//byte casting
	for l := byte('a'); l <= byte('z'); l++ {
		//casting back to a string
		go println(string(l))
	}
}
