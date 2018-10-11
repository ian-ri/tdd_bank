package main

type account struct {

}


func NewAccount (amount int64) *account  {


	if (amount < 0) {
		return nil
	}

	return &account{}

}