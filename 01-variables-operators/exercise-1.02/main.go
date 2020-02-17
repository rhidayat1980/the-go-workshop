package main

import "fmt"

var (
	firstName     string = "rachmat"
	lastName      string = "hidayat"
	age           int    = 39
	peanutAllergy bool   = false
)

func main() {
	// accessing variable
	var foo string = "bar"
	var bas string = "qux"
	fmt.Println(foo, bas)

	fmt.Println("accessing variable declared with var ()")
	fmt.Println("firstname: ", firstName)
	fmt.Println("lastName: ", lastName)
	fmt.Println("age: ", age)
	fmt.Println("peanutAllergy: ", peanutAllergy)
}
