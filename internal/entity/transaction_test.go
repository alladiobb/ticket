package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane  2", "jane@j.com")
	account2 := NewAccount(client2)

	account1.Credit(1000.0)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 500.0)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 500.0, account1.Balance)
	assert.Equal(t, 1500.0, account2.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane  2", "jane@j.com")
	account2 := NewAccount(client2)

	account1.Credit(1000.0)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 2000.0)
	assert.NotNil(t, err)
	assert.Error(t, err, "insuficient balance")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}
