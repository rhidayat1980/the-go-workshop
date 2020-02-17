package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	debug       bool
	logLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(debug, logLevel, startUpTime)

	var seed int64 = 1234456789

	rand.Seed(seed) // seed need int64
}
