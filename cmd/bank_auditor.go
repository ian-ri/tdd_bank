package main

//BankAuditor is an interface
type BankAuditor interface {
	OverdraftAlert(amount int64)
}
