package account

type Account struct {
	balance int64
}

func (a *Account) CheckBalance() int64 {
	return a.balance
}

func NewAccount (amount int64) *Account {


	if (amount < 0) {
		return nil
	}


	return &Account{balance: amount}

}

func  (a *Account) Withdraw(amount int64)  {
	a.balance= a.balance-amount
}