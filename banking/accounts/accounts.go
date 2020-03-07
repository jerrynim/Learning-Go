package accounts

import "errors"

//Account strict
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Banlance of your account
func (a Account) Balance() int {
	return a.balance
}

// withdraw x amount from your account
func (a Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("cant with draw")
	}
	a.balance -= amount
	return nil
}

func (a Account) String() string {
	return "it is Account struct"
}
