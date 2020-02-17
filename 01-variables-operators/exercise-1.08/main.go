// changing multiple values at once

package main

import "fmt"

func main() {
	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)

	fmt.Println("\nchanging multiple values")
	query, limit, offset = "bar", 20, 1
	fmt.Println(query, limit, offset)
}
