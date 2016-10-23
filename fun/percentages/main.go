package main

import "fmt"

func main() {
	originalPrice := 20
	discountPercentage := .30
	PriceToFloat := float64(originalPrice)
	discountValue := PriceToFloat * discountPercentage
	fmt.Println("The discounted price is:", discountValue)
}
