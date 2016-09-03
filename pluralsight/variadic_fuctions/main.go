package main

import (
	"fmt"
)

func main() {
	bestFinish := bestLeagueFinishes(14, 10, 13, 17, 14, 16, 2, 1)
	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {

	//this is 13, first item in the slice
	best := finishes[0]
	fmt.Println("Best is:", best)

	//loop through slice and if slice is < best assign best
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
