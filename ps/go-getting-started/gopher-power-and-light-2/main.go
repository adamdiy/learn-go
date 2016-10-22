package main

import (
	"fmt"
)

func main() {


	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("1) Generate Power Grid Report")
	fmt.Println("Please choose an option: ")

	var option string

	//& = pass memory address location of option to Scanln
	fmt.Scanln(&option)
	//println(option)

	plantCapacities := []float64{ 30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.

	switch option {
	case "1":
		for idx, cap := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
		}
	case "2":
		capacity := 0.
		for _, plantId := range activePlants {
			capacity += plantCapacities[plantId]
		}

	utilization := gridLoad / capacity
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization", utilization * 100)

	default:
		fmt.Println("Unknown option, no action taken")
	}
}
