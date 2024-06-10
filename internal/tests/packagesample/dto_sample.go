package packagesample

import (
	"github.com/guregu/null/v5"
	"github.com/saucon/sulaptomap/internal/tests/packagesampleother"
)

type SomeStructA struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`

	Age         int     `json:"age" gator:"gte=0,lte=100"`
	TotalAmount float64 `json:"totalAmount"`

	IsActive bool `json:"isActive"`

	WhatTheFuckIsComplex complex64 `json:"whatTheFuckIsComplex"`

	AmountCurrency AmountCurrency
	User           packagesampleother.User

	Amount   *float64 `json:"amount"`
	Currency *string  `json:"currency"`

	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`

	CreatedAt null.Time `json:"createdAt"`
	UpdatedAt null.Time `json:"updatedAt"`
}

func (abc *SomeStructA) A() {}

type AmountCurrency struct {
	Amount   float64
	Currency string
}

type SomeStructB struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`

	Amount   *float64 `json:"amount"`
	Currency *string  `json:"currency"`

	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`

	CreatedAt null.Time `json:"createdAt"`
	UpdatedAt null.Time `json:"updatedAt"`
}
