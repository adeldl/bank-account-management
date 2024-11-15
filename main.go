package main

import (
	"bank-account-management/bank"
	"fmt"
)

func main() {
	account := bank.NewAccount() // Gunakan konstruktor

	// Deposit money
	if err := account.Deposit(100.50); err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to withdraw money
	if err := account.Withdraw(50); err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to withdraw more than the balance
	if err := account.Withdraw(100); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Balance:", account.GetBalance())
}
