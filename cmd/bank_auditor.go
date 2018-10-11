package main

type BankAuditor interface {
	OverdraftAlert(amount int64)
}
