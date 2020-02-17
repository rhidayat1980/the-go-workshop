package main

import "fmt"

func main() {
	names := []string{
		"jim",
		"Jane",
		"Joe",
		"June",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for k, v := range names {
		fmt.Println(k, v)
	}

	for _, v := range names {
		fmt.Println(v)
	}

	for k := range names {
		fmt.Println(k)
	}

}
