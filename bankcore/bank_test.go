package bank

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angels, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angels, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angels, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angels, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("Statement doesn't have the proper format")
	}
}

func TestTransfer(t *testing.T) {
	account1 := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angels, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account2 := Account{
		Customer: Customer{
			Name:    "Doe",
			Address: "Shanghai, China",
			Phone:   "(+86) 177 9871 3781",
		},
		Number:  1002,
		Balance: 0,
	}

	account1.Deposit(100)
	err := account1.Transfer(&account2, 50)
	if account1.Balance != 50 || account2.Balance != 50 {
		t.Error("Transfer not work", err)
	}
}
