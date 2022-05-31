package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrediCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Luiz Reginaldo", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Luiz Reginaldo", 12, 2024, 123)
	assert.Nil(t, err)
}