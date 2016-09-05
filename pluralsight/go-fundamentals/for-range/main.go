package main

import (
	"fmt"
)

func main() {
	courseInProg := []string{"Docker Deep Dive", "Docker Clustering", 
	"Docker and Kuberetes"}
	courseCompleted := []string{"Puppet essentials", "Golang in action", "Docker Deep Dive"}
	//fmt.Println(courseInProg[1]) }
	for _, i := range courseInProg {
		fmt.Println(i)
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("\nHey we found a clash", i, "is in both lists")
			}
		}
	}
}
