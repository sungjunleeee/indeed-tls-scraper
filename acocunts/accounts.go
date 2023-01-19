package accounts

// Account struct
type Account struct {
	// Public field: Uppercase, Private field: Lowercase
	owner   string
	balance int
}

// NewAccount creates Account
// *Account means return a pointer to Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // return a copy (memory address) of account
}
