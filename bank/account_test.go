package bank_test

import (
	"bank-account-management/bank"
	"testing"
)

func TestDeposit(t *testing.T) {
	account := bank.NewAccount() // Menggunakan konstruktor untuk membuat akun baru
	err := account.Deposit(100)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if account.GetBalance() != 100 {
		t.Errorf("Expected balance of 100, got %v", account.GetBalance())
	}
}

func TestWithdraw(t *testing.T) {
	account := bank.NewAccount()
	account.Deposit(100) // Deposit awal

	err := account.Withdraw(50)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if account.GetBalance() != 50 {
		t.Errorf("Expected balance of 50, got %v", account.GetBalance())
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	account := bank.NewAccount()
	account.Deposit(100) // Deposit awal

	err := account.Withdraw(150) // Mencoba menarik lebih dari saldo
	if err == nil {
		t.Error("Expected error for insufficient funds, got none")
	}
	if account.GetBalance() != 100 {
		t.Errorf("Expected balance to remain 100, got %v", account.GetBalance())
	}
}

func TestDepositNegativeAmount(t *testing.T) {
	account := bank.NewAccount()
	err := account.Deposit(-50) // Mencoba deposit jumlah negatif
	if err == nil {
		t.Error("Expected error for negative deposit, got none")
	}
	if account.GetBalance() != 0 {
		t.Errorf("Expected balance to remain 0, got %v", account.GetBalance())
	}
}
