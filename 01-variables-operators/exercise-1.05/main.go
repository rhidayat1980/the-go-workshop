package main

import (
	"fmt"
	"time"
)

func main() {
	debug := false
	logLevel := "info"
	startUpTime := time.Now()

	fmt.Println("debug mode:", debug, "\nlog Level:", logLevel, "\nstart at:", startUpTime)

	// declaring multiple variable with short declaration
	firstName, lastName, age, bitcoinUser := "rachmat", "hidayat", 39, true
	fmt.Println("firstName:", firstName,
		"\nlastName:", lastName, "\nage:", age, "\nbitcoinUser:", bitcoinUser)
}
