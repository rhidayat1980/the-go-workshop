package main

import (
	"errors"
	"fmt"
)

func main() {
	input := 21

	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}

func validate(input int) error {
	if input < 0 {
		return errors.New("input can't be negative number")
	} else if input > 100 {
		return errors.New("input can't be greater than 100")
	} else if input%7 == 0 {
		return errors.New("input can not be divisible by 7")
	} else {
		return nil
	}

}
