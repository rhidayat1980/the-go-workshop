package main

import "fmt"

// const (
// 	sunday    = 0
// 	monday    = 1
// 	tuesday   = 2
// 	wednesday = 3
// 	thursday  = 4
// 	friday    = 5
// 	saturday  = 6
// )

const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	fmt.Println("sunday:", sunday)
	fmt.Println("monday:", monday)
	fmt.Println("tuesday:", tuesday)
	fmt.Println("wednesday:", wednesday)
	fmt.Println("thursday:", thursday)
	fmt.Println("friday:", friday)
	fmt.Println("saturday", saturday)
}
