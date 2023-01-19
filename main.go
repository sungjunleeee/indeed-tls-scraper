package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	// one way of doing for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	// another way of doing for loop
	for _, number := range numbers { // you can simply use number instead of _, number
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
