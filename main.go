package main

import "fmt"

func main() {
	// [key type]value type
	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
}
