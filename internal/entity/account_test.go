package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(c)
	assert.NotNil(t, account)
	assert.Equal(t, c.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j")
	account := NewAccount(c)
	account.Credit(100)
	assert.Equal(t, 100.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j")
	account := NewAccount(c)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}
