package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, c)
	// assert.Equal(t, "1", c.ID)
	assert.Equal(t, "John Doe", c.Name)
	assert.Equal(t, "j@j.com", c.Email)
}

func TestNewClientWhenArgsInvalid(t *testing.T) {
	c, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, c)
}

func TestUpdateClient(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j.com")
	err := c.Update("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", c.Name)
	assert.Equal(t, "j@j.com", c.Email)
}

func TestUpdataClienteWithInvalidArgs(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j.com")
	err := c.Update("", "j@j.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(c)
	err := c.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.Accounts))
}
