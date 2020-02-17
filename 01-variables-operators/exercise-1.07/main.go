package main

import "fmt"

func main() {
	offset := 10
	fmt.Println(offset)

	offset = 5
	fmt.Println(offset)

	defaultOffset := 20
	fmt.Println(defaultOffset)

	defaultOffset = offset + defaultOffset
	fmt.Println(defaultOffset)

}
