// declaring multiple var from function
package main

import (
	"fmt"
	"time"
)

func main() {
	debug, logLevel, startUpTime := getConfig()
	fmt.Println(debug, logLevel, startUpTime)
}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
