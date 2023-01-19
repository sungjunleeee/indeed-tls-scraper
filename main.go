package main

import (
	"fmt"

	"github.com/sungjunleeee/learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	// Interestingly enough
	// a short variable declaration can legal
	// only when you declare multi-varible short declaration (at least one varible should be new)
	err = dictionary.Delete(baseWord)
}
