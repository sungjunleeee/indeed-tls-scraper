package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	// Public field: Uppercase, Private field: Lowercase
	owner   string
	balance int
}

// Error starts with err
// you cannot use short variable declaration on the top level
// (this simpifies parsing)
var errNoMoney = errors.New("can't withdraw")

// NewAccount creates Account
// *Account means return a pointer to Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // return a copy (memory address) of account
}

// Deposit x amount on your account
// (a *Account) is called a receiver in Go
// a can be anything, but it's a convention to start with first lowercase letter of the struct
// *Account means a pointer to Account, not a copy
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s acount.\nHas: ", a.Balance())
}
