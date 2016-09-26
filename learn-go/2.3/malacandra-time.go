package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		//time to arrive * 24 hours
		tripDuration = 28 * (24 * time.Hour)
		tripInHours = 28 * 24
		tripInSeconds = 28 * 86400
		distanceToMalacandra = 56000000 //km/s
	)

	// Solve for unknown speed in km/s
	// We determine seconds
	// secondsToArrival := timeToArriveInDays * secondsInDay
	fmt.Println(tripDuration)
	fmt.Println(tripInHours)
	fmt.Println(tripDuration.Seconds())
	fmt.Println(tripInSeconds)
	fmt.Println(tripDuration.Minutes())
	fmt.Println(tripDuration.Hours())
}
