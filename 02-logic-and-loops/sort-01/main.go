package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{4, 3, 8, 1, 2, 5, 7, 6, 9, 11, 10}
	fmt.Println("before sorting", s)

	sort.Ints(s)
	fmt.Println("after sorting", s)
}
