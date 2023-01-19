package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"chicken", "pizza"}
	jun := person{name: "jun", age: 25, favFood: favFood}
	fmt.Println(jun.favFood)
}
