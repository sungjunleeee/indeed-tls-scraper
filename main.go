package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// an initialization statement to set up a local variable
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
