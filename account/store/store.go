package store

import "errors"

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Account
	Credits float64
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	// 没有断言
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return 0.0, nil
		} else {
			return 0.0, errors.New("你不能删除超过账户的数值")
		}
	} else {
		return 0.0, errors.New("不能以负值作输入")
	}
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0 {
		e.Credits += amount
		return 0.0, nil
	} else {
		return 0.0, errors.New("不能以负值作输入")
	}
}

func (a *Account) ChangeName(name string) {
	a.FirstName = name
}
