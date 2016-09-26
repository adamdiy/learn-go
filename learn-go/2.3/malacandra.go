package main

import (
	"fmt"
)

func main() {
	const (
		timeToArriveInDays = 28
		secondsInDay = 86400

		distanceToMalacandra = 56000000
	)

	// unknown = speed in km/s

	totalSeconds := timeToArriveInDays * secondsInDay
	fmt.Println("Total amount of seconds: ",totalSeconds)

	speed := distanceToMalacandra / totalSeconds
	fmt.Println("Speed required: ", speed) // 23 km/s?


	//check answer:
	fastestSpaceTravelNow := distanceToMalacandra / speed
	fmt.Println("Total amount of seconds: ", fastestSpaceTravelNow)
	
}
