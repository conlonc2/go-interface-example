package main

import (
	"fmt"

	"github.com/conlonc2/gochain/pkg/account"
)

func runDefault() {
	defaultAccount := account.NewDefaultAccount()

	balance := defaultAccount.GetBalance()
	afterDeposit := defaultAccount.Deposit(5)
	afterWithdraw, _ := defaultAccount.Withdraw(2)
	afterWithdraw2, errExpected := defaultAccount.Withdraw(4)

	fmt.Printf("Default Account: Initial Balance: %v\nAfter Deposit: %v\nAfter First Withdraw: %v\nAfter Second Withdraw: %v\n", balance, afterDeposit, afterWithdraw, afterWithdraw2)

	if errExpected != nil {
		fmt.Println(errExpected)
	}
}

func runA() {
	a := account.NewAAccount()
	balance := a.GetBalance()
	afterDeposit := a.Deposit(5)
	afterWithdraw, _ := a.Withdraw(2)
	afterWithdraw2, errExpected := a.Withdraw(4)

	fmt.Printf("A Account: Initial Balance: %v\nAfter Deposit: %v\nAfter First Withdraw: %v\nAfter Second Withdraw: %v\n", balance, afterDeposit, afterWithdraw, afterWithdraw2)

	if errExpected != nil {
		fmt.Println(errExpected)
	}
}

func runB() {
	b := account.NewAAccount()
	balance := b.GetBalance()
	afterDeposit := b.Deposit(5)
	afterWithdraw, _ := b.Withdraw(2)
	afterWithdraw2, errExpected := b.Withdraw(4)

	fmt.Printf("B Account: Initial Balance: %v\nAfter Deposit: %v\nAfter First Withdraw: %v\nAfter Second Withdraw: %v\n", balance, afterDeposit, afterWithdraw, afterWithdraw2)

	if errExpected != nil {
		fmt.Println(errExpected)
	}
}

func main() {
	fmt.Println("----------------------")
	runDefault()
	fmt.Println("----------------------")
	runA()
	fmt.Println("----------------------")
	runB()
}
