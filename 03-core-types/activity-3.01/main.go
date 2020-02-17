package main

import "fmt"

func main() {
	taxTotal := .0

	// cake
	taxTotal += salesTax(.99, .075)

	// milk
	taxTotal += salesTax(2.75, .015)

	// butter
	taxTotal += salesTax(.87, .02)

	// total
	fmt.Println("Sales Tax Total: ", taxTotal)
}

func salesTax(cost, taxRate float64) float64 {
	return cost * taxRate
}
