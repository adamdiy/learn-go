package main

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)


	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()
	//The () = self executing function

	go func() {
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()

	//time.Sleep(6 * time.Second)
}
