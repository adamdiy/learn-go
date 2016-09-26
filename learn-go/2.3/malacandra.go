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

	// Solve for unknown speed in km/s
	// We determine seconds
	secondsToArrival := timeToArriveInDays * secondsInDay
	minsToArrival := secondsToArrival / 60
	hrsToArrival := minsToArrival / 60
	fmt.Println("Total amount of seconds to arrival:", secondsToArrival)
	fmt.Println("Total amount of mins to arrival:", minsToArrival)
	fmt.Println("Total amount of hours to arrival:", hrsToArrival)

	// formula for speed = distance / time
	// speed := distanceToMalacandra / secondsToArrival
	// since speed is measured in seconds we divide distance / seconsToArrival
	speed := distanceToMalacandra / secondsToArrival
	fmt.Println("Speed required to arrive in 28 days:", speed, "km/s") // 23 km/s?

	//check answer:
	fastestSpaceTravelNow := distanceToMalacandra / speed
	fmt.Println("Total amount of seconds: ", fastestSpaceTravelNow)
	
}
