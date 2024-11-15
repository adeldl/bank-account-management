package bank

import (
	"fmt"
)

type Account struct {
	balance float64
}

func NewAccount() *Account {
	return &Account{
		balance: 0,
	}
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > a.balance {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}
