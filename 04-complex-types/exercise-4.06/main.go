// looping over array

package main

import "fmt"

func main() {
	fmt.Println(message())
}

func message() string {
	m := ""
	arr := []int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m

}
