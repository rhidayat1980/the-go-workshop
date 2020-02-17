package main

import "fmt"

func main() {

	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(rune(runeVal)))
	}

	name := "John Doe"

	for k, v := range name {
		fmt.Println(k, v)
	}

	for k, v := range string(name) {
		fmt.Println(k, v)
	}

	for k, v := range name {
		fmt.Println(k, string(rune(v)))
	}
}
