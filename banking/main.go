package main

import (
	"errors"
	"fmt"

	"github.com/jerrynim/learngo/banking/accounts"
)

var err_noMoney = errors.New("cant")

func main() {
	account := accounts.NewAccount("jerrynim")
	account.Deposit(10)
	fmt.Println(account)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err_noMoney)
	}
	fmt.Println(account.Balance())
}
