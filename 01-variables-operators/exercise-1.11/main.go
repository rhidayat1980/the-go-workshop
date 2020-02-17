package main

import "fmt"

func main() {
	visits := 15

	fmt.Println("first visits		:", visits == 1)
	fmt.Println("return visits		:", visits != 1)
	fmt.Println("silver members		:", visits >= 10 && visits < 21)
	fmt.Println("gold members		:", visits > 20 && visits < 30)
	fmt.Println("platinum members	:", visits > 30)
}
