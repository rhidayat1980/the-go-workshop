package main

import "fmt"

func main() {

	var count int

	add5value(count)
	fmt.Println("add5value post :", count)

	add5point(&count)
	fmt.Println("add5point post", count)
}

func add5value(count int) {
	count += 5
	fmt.Println("add5value :", count)
}

func add5point(count *int) {
	*count += 5
	fmt.Println("add5point :", *count)
}
