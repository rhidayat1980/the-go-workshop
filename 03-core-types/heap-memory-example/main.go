package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("with int")
	var list []int

	for i := 0; i < 10_000_000; i++ {
		list = append(list, 100)
	}

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)

	fmt.Println("with int8")

	var list2 []int8

	for i := 0; i < 10_000_000; i++ {
		list2 = append(list2, 100)
	}

	var n runtime.MemStats

	runtime.ReadMemStats(&n)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", n.TotalAlloc/1024/1024)
}
