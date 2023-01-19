package main

import (
	"fmt"

	accounts "github.com/sungjunleeee/learngo/acocunts"
)

func main() {
	account := accounts.NewAccount("jun")
	account.Deposit(500)
	fmt.Println(account.Balance())
	err := account.Withdraw(600)
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err) // fmt.Println(err) + os.Exit(1)
	}
	fmt.Println(account.Balance())
}
