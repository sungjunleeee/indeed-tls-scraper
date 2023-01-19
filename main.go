package main

import (
	"fmt"
	"time"
)

func main() {
	go sickCount("Jun")
	sickCount("Tim")
}

func sickCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is super sick 🥵", i)
		time.Sleep(time.Second)
	}
}
