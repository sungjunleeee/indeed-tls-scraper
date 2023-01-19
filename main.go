package main

import (
	"fmt"

	accounts "github.com/sungjunleeee/learngo/acocunts"
)

func main() {
	account := accounts.NewAccount("jun")
	fmt.Println(account)
}
