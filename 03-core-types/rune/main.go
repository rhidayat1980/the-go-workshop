package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}

	fmt.Println("")

	for k, v := range username {
		fmt.Println(k, " = ", v)
	}

	fmt.Println("")

	runes := []rune(username)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}

	fmt.Println("")

	for _, v := range runes {
		fmt.Println(string(v))
	}
}
