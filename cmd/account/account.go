package account

type Account struct {
	name string
	balance int64
}

func (a *Account) CheckBalance() int64 {
	return a.balance
}

func (a *Account) GetName() string {
	return a.name
}

func NewAccount (name string, amount int64) *Account {


	if (amount < 0) {
		return nil
	}


	return &Account{name: name,balance: amount}

}

func  (a *Account) Withdraw(amount int64)  {
	a.balance= a.balance-amount
}