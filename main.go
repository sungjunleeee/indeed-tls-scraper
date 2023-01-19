package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer is executed after the function is done
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // you still need return keyword
}

func main() {
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)
}
