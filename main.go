package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool) // initialize channel
	people := [2]string{"Jun", "Tim"}
	for _, person := range people {
		// run go routine with channel
		go isSick(person, c)
	}
	// wait for the channel to return a value
	fmt.Println(<-c)
	fmt.Println(<-c)
	// result := <-c
	// fmt.Println(result)
}

func isSick(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	// return value to channel
	c <- true
}
