package account

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	t.Run( "is account created", func(t *testing.T) {
		account := NewAccount("some-one",10)
		require.NotNil(t,account)
	} )

	t.Run( "Account not created for negative amount", func(t *testing.T) {
		account := NewAccount("some-one",-10)
		require.Nil(t,account)
	} )

	t.Run("check initial balance", func(t *testing.T) {
		account := NewAccount("some-one",50)
		require.Equal(t, int64(50), account.CheckBalance())
	})

	t.Run("check account name", func(t *testing.T) {
		account := NewAccount("some-one",50)
		require.Equal(t, "some-one",account.GetName())
	})

	t.Run("withdraw money", func(t *testing.T) {
		account := NewAccount("some-one",50)
		account.Withdraw(10)
		require.Equal(t, int64(40), account.CheckBalance())
	})

}

