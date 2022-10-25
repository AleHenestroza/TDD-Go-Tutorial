package main

import "testing"

func TestBadBank(t *testing.T) {
	var (
		paola  = Account{Name: "Paola", Balance: 100}
		juan   = Account{Name: "Juan", Balance: 75}
		susana = Account{Name: "Susana", Balance: 200}

		transactions = []Transaction{
			NewTransaction(juan, paola, 100),
			NewTransaction(susana, juan, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(paola), 200)
	AssertEqual(t, newBalanceFor(juan), 0)
	AssertEqual(t, newBalanceFor(susana), 175)

}
