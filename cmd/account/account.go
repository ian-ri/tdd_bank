package account

type Account struct {

}


func NewAccount (amount int64) *Account {


	if (amount < 0) {
		return nil
	}

	return &Account{}

}