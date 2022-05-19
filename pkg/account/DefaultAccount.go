package account

import (
	"errors"
	"fmt"
	"math"
)

type IAccount interface {
	GetBalance() float64
	Withdraw(amount float64) (float64, error)
	Deposit(amount float64) float64
}

type DefaultAccount struct {
	balance float64
}

func NewDefaultAccount() *DefaultAccount {
	return &DefaultAccount{balance: 0}
}

func (a *DefaultAccount) GetBalance() float64 {
	return math.Round(a.balance)
}

func (a *DefaultAccount) Withdraw(amount float64) (float64, error) {
	newBalance := a.balance - amount
	if newBalance < 0 {
		errMsg := fmt.Sprintf("Invalid account balance: Attempted to withdraw %v and current balance is: %v", amount, a.GetBalance())
		return a.balance, errors.New(errMsg)
	}
	a.balance = newBalance
	return newBalance, nil
}

func (a *DefaultAccount) Deposit(amount float64) float64 {
	a.balance += amount
	return a.balance
}
