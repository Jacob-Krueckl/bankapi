package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/msft/bank"
)

var accounts = map[float64]*bank.Account{}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number: 1001,
	}

	accounts[1002] = &bank.Account{
		Customer: bank.Customer{
			Name:    "Emily",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number: 1002,
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {

	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing account number")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "invalid account number")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "account with number %v not found", number)
		} else {
			json.NewEncoder(w).Encode(bank.Statement(account))
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" || amountqs == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing account number or amount")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "invalid account number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "invalid amount")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "account with number %v not found", number)
		} else {
			if err := account.Deposit(amount); err != nil {
				fmt.Fprintf(w, "error depositing: %v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" || amountqs == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing account number or amount")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "invalid account number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "invalid amount")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "account with number %v not found", number)
		} else {
			if err := account.Withdraw(amount); err != nil {
				fmt.Fprintf(w, "error withdrawing: %v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	fromqs := req.URL.Query().Get("from")
	toqs := req.URL.Query().Get("to")
	amountqs := req.URL.Query().Get("amount")

	if fromqs == "" || toqs == "" || amountqs == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing from account, to account, or amount")
		return
	}

	if from, err := strconv.ParseFloat(fromqs, 64); err != nil {
		fmt.Fprintf(w, "invalid from account number")
	} else if to, err := strconv.ParseFloat(toqs, 64); err != nil {
		fmt.Fprintf(w, "invalid to account number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "invalid amount")
	} else {
		fromAccount, ok := accounts[from]
		if !ok {
			fmt.Fprintf(w, "account with number %v not found", from)
		} else {
			toAccount, ok := accounts[to]
			if !ok {
				fmt.Fprintf(w, "account with number %v not found", to)
			} else {
				if err := fromAccount.Transfer(toAccount, amount); err != nil {
					fmt.Fprintf(w, "error transferring: %v", err)
				} else {
					fmt.Fprintf(w, fromAccount.Statement())
					fmt.Fprintf(w, "\n")
					fmt.Fprintf(w, toAccount.Statement())
				}
			}
		}
	}
}
