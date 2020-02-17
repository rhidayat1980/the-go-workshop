package main

import "fmt"

func main() {
	// main course
	var total float64 = 2 * 13
	fmt.Println("Sub:", total)

	// drinks
	total = total + (4 * 2.25)
	fmt.Println("Sub:", total)

	// discount
	total = total - 5
	fmt.Println("Sub:", total)

	// 10% tip
	tip := total * 0.1
	fmt.Println("Tip:", tip)

	// finally we add the tip to the total
	total = total + tip
	fmt.Println("Total:", total)

	// split bill
	split := total / 2
	fmt.Println("Split:", split)

	// reward every 5th visit
	visitCount := 24
	visitCount = visitCount + 1

	remainder := visitCount % 5

	if remainder == 0 {
		fmt.Println("With the visit, you have earned a reward")
	}

	givenName := "John"
	familyName := "Smith"
	fullName := givenName + " " + familyName
	fmt.Println("Hello, ", fullName)
}
