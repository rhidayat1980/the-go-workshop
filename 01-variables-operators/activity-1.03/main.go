package main

import "fmt"

func main() {
	count := 5
	message := ""
	if count > 5 {
		message = "greater than 5"
	} else {
		message = "not greater than 5"
	}

	fmt.Println("message", message)
}
