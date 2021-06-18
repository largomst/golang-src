package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	} else {
		a.Balance += amount
		return nil
	}
}
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	} else if a.Balance < amount {
		return errors.New("the amount to withdraw should be less than the account's balance")
	} else {
		a.Balance -= amount
		return nil
	}
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%d - %s - %v", a.Number, a.Name, a.Balance)

}
func Hello() string {
	return "Hey, I'm working!"
}

func (a *Account) Transfer(b *Account, amount float64) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	} else {
		if err := b.Deposit(amount); err != nil {
			return err
		} else {
			return nil
		}
	}
}
