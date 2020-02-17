package main

import "fmt"

func main() {

	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(rune(runeVal)))
	}

	fmt.Println("")
	name := "John Doe"

	for k, v := range name {
		fmt.Println(k, v)
	}
	fmt.Println("")
	for k, v := range string(name) {
		fmt.Println(k, v)
	}

	fmt.Println("")
	for key, value := range name {
		fmt.Println(key, string(rune(value)))
	}
	fmt.Println("")
	for k, v := range name {
		fmt.Println(k, string(rune(v)))
	}
}
