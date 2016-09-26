package main

import "fmt"

func main() {
	const (
		fastestCurrentSpeed = 16 // km/s //fastest space ships only travel 16 km/s
		lightSpeed = 299792 // km/s
		secondsInMinute = 60
		minutesInHour = 60
		hoursInDay = 24
	)
	//var distance = 56000000 // km
	var distance = 57600000 //km //distance on July 18, 2018

	fastestSpaceTravelNow := distance / fastestCurrentSpeed
	timeInMinutesNow := fastestSpaceTravelNow / secondsInMinute
	timeInHoursNow := timeInMinutesNow / minutesInHour
	timeInDaysNow := timeInHoursNow / hoursInDay

	timeInSecondsWithLight := distance / lightSpeed
	timeInMinutesWithLight := timeInSecondsWithLight / secondsInMinute
	timeInHoursWithLight := timeInMinutesWithLight / minutesInHour

	fmt.Println("Time in seconds w/ current technology: ", fastestSpaceTravelNow)
	fmt.Println("Time in minutes w/ current technology:", timeInMinutesNow)
	fmt.Println("Time in hours w/ current technology:", timeInHoursNow)
	fmt.Println("Time in days w/ current technology:", timeInDaysNow)
	fmt.Println("")
	fmt.Println("Time in seconds w/ light technology:", timeInSecondsWithLight)
	fmt.Println("Time in minutes w/ light technology:", timeInMinutesWithLight)
	fmt.Println("Time in hours w/ light technology:", timeInHoursWithLight)

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}
