package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")

	if order.Validate() == nil{
		t.Error("ID is required")
	}
}


func TestIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "1"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestFinalPrice(t *testing.T){
	order := Order{ID: "2", Price: 10, Tax: 1}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "2", order.ID)
	assert.Equal(t, 10, order.Price)
}