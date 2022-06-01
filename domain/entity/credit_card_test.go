package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCrediCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Luiz Reginaldo", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Luiz Reginaldo", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Luiz Reginaldo", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Luiz Reginaldo", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Luiz Reginaldo", 11, 2024, 123)
	assert.Nil(t, err)
}

func TestCredidCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("4193523830170205", "Luiz Reginaldo", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
