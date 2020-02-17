// Activity 4.02: Printing a User's Name Based on User Input

package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	name, exists := getName(os.Args[1])
	if !exists {
		fmt.Printf("error: user (%v) not found", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Hi,", name)
}

func getName(id string) (string, bool) {
	name, exist := users[id]
	return name, exist
}
