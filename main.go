package main

import (
	"fmt"

	"github.com/sungjunleeee/learngo/dict"
)

func main() {
	// dict is not a reserved keyword
	dictionary := dict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("Found:", word, "Definition:", hello)
	// generating error
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
