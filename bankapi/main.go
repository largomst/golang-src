package main

import (
	"bank"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*bank.Account{}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angles, California",
			Phone:   "(213) 555 0417",
		}, Number: 1001,
	}
	accounts[1002] = &bank.Account{
		Customer: bank.Customer{
			Name:    "Doe",
			Address: "Shanghai, China",
			Phone:   "(+86) 177 9871 3781",
		},
		Number:  1002,
		Balance: 0,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/desposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprint(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")
	if numberqs == "" {
		fmt.Fprint(w, "Account number is missing!")
		return
	} else {
		if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
			fmt.Fprint(w, "Invalid account number")
		} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
			fmt.Fprint(w, "Invalid amount number")
		} else {
			account, ok := accounts[number]
			if !ok {
				fmt.Fprintf(w, "Account with number %v can't be found!", number)
			} else {
				err := account.Deposit(amount)
				if err != nil {
					fmt.Fprintf(w, "%v", err)
				} else {
					fmt.Fprint(w, account.Statement())
				}
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	} else {
		if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
			fmt.Fprintf(w, "Invalid account number!")
		} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
			fmt.Fprintf(w, "Invalid amount number!")
		} else {
			account, ok := accounts[number]
			if !ok {
				fmt.Fprintf(w, "Account with number %v can't be found!", number)
			} else {
				err := account.Withdraw(amount)
				if err != nil {
					fmt.Fprintf(w, "%v", err)
				} else {
					fmt.Fprint(w, account.Statement())
				}
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	number := req.URL.Query().Get("number")
	destqs := req.URL.Query().Get("destqs")
	amountqs := req.URL.Query().Get("amount")
	if number == "" || destqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	} else {
		if number_from, err := strconv.ParseFloat(number, 64); err != nil {
			fmt.Fprintf(w, "Invalid source account number")
		} else if number_to, err := strconv.ParseFloat(destqs, 64); err != nil {
			fmt.Fprintf(w, "Invalid target account number")
		} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
			fmt.Fprintf(w, "Invalid amount number!")
		} else {
			if accountA, ok := accounts[number_from]; !ok {
				fmt.Fprintf(w, "Account with number %v can't be found!", number_from)
			} else if accountB, ok := accounts[number_to]; !ok {
				fmt.Fprintf(w, "Account with number %v can't be found!", number_to)
			} else {
				err := accountA.Transfer(accountB, amount)
				if err != nil {
					fmt.Fprintf(w, "%v", err)
				} else {
					fmt.Fprint(w, accountB.Statement())
				}
			}

		}
	}

}
