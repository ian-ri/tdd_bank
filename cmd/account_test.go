package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	t.Run( "is account created", func(t *testing.T) {
		account := NewAccount(10)
		require.NotNil(t,account)
	} )

	t.Run( "Account not created for negative amount", func(t *testing.T) {
		account := NewAccount(-10)
		require.Nil(t,account)
	} )
}

