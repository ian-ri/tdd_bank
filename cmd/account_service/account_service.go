package account_service

import "errors"

type accountService struct {
	balance int64
	accountName string
}

type AccountService interface {
	Open(accountName string, initialAmount int64) error
	CheckBalance(accountName string) (int64,error)

}

func (a *accountService) Open(accountName string, initialAmount int64) error {

	if initialAmount < 0 {
		return errors.New("Negative Initial Amount is not permitted")
	}
	a.accountName = accountName
	a.balance = initialAmount
	 return nil
}

func NewAccountService() AccountService {
	return &accountService{}
}

func (a *accountService) CheckBalance(accountName string) (int64,error) {

	if a.accountName != accountName {
		return 0, errors.New("Account does not exist")
	}
	return a.balance, nil
}


