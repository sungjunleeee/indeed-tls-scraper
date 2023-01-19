package main

import (
	"fmt"

	"github.com/sungjunleeee/learngo/something"
)

func main() {
	fmt.Println("Hello, World!")
	something.SayHello()
	something.sayBye() // This will not work because sayBye() is not exported
}
