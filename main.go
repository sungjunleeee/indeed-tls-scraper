package main

import (
	"fmt"

	accounts "github.com/sungjunleeee/learngo/acocunts"
)

func main() {
	account := accounts.NewAccount("jun")
	account.Deposit(500)
	fmt.Println(account)
}
