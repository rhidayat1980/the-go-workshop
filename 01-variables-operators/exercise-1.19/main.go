package main

import "fmt"

var level = "pkg"

func main() {
	fmt.Println("main start:", level)

	// create shadow variable
	level := 42

	if true {
		fmt.Println("block start:", level)
		funcA()
	}

	fmt.Println("main end:", level)
}

func funcA() {
	fmt.Println("funcA start:", level)
}
