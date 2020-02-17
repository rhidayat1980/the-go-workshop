package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, World",
	"Καλημέρα κόσμε",

	"こんにちは世界",

	"سلام دنیا‎",

	"Привет, мир",
}

func main() {

	// seed random number generator from the current time
	rand.Seed(time.Now().UnixNano())

	// generate random number in the range of out list
	index := rand.Intn(len(helloList))

	// call a function and receive multiple return values
	msg, err := hello(index)

	// handle any errors
	if err != nil {
		log.Fatal(err)
	}

	// printout message to console
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {

		// create an error, convert the int type to a string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}
