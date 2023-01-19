package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// you can also use initialization statement in switch like if else statement
	// one way of using switch
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	// another way of using switch
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
