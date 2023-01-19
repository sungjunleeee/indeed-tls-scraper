package main

import "fmt"

func main() {
	array := [5]string{"nico", "lynn", "dal"}
	slice := []string{"nico", "lynn", "dal"}
	slice = append(slice, "flynn") // new slice is returned from append
	fmt.Println(slice)
	fmt.Println(array)
}
