package main

import (
	"fmt"

	"github.com/sungjunleeee/learngo/dict"
)

func main() {
	// dict is not a reserved keyword
	dictionary := dict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
