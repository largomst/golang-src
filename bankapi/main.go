package main

import (
	"bank"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]bank.Account{}

func main() {
	accounts[1001] = bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angles, California",
			Phone:   "(213) 555 0417",
		}, Number: 1001,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/desposit", deposit)
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
