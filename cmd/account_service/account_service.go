package account_service

import (
	"errors"
	"github.com/mmircea16/tdd_bank/cmd/account"
)

type accountService struct {
	accounts map[string]*account.Account
}

type AccountService interface {
	Open(accountName string, initialAmount int64) error
	CheckBalance(accountName string) (int64,error)
	Withdraw(accountName string, amount int64) error
}



func (a *accountService) Open(accountName string, initialAmount int64) error {

	if initialAmount < 0 {
		return errors.New("Negative Initial Amount is not permitted")
	}
	a.accounts[accountName] = account.NewAccount(accountName,initialAmount)
	return nil
}

func NewAccountService() AccountService {
	myAccounts := make(map[string]*account.Account)
	return &accountService{accounts:myAccounts}

}

func (a *accountService) accountExists(accountName string) bool {
	if _,ok := a.accounts[accountName] ; !ok {
		return false
	}
	return true
}

func (a *accountService) CheckBalance(accountName string) (int64,error) {

	if !a.accountExists(accountName) {
		return 0, errors.New("Account does not exist")
	}
	return a.accounts[accountName].CheckBalance(), nil
}

func (a *accountService) Withdraw(accountName string, amount int64) error {

	if !a.accountExists(accountName) {
		return errors.New("Account does not exist")
	}
	a.accounts[accountName].Withdraw(amount)
	return nil
}



