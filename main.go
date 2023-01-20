package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // initialize channel
	people := [5]string{"Jun", "Tim", "Comico", "Nico", "Flynn"}
	for _, person := range people {
		// run go routine with channel
		go isSick(person, c)
	}
	fmt.Println("Waiting for messages")
	for _, person := range people {
		go isSick(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
	// result := <-c
	// fmt.Println(result)
}

func isSick(person string, c chan string) {
	time.Sleep(time.Second * 5)
	// return value to channel
	c <- person + " is sick 🥵"
}
