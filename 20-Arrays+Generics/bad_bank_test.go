package main

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Juan",
			To:   "Paola",
			Sum:  100,
		},
		{
			From: "Susana",
			To:   "Juan",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Paola"), 100)
	AssertEqual(t, BalanceFor(transactions, "Juan"), -75)
	AssertEqual(t, BalanceFor(transactions, "Susana"), -25)
}
