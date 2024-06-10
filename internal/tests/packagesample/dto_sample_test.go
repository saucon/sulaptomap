package packagesample

import (
	"github.com/guregu/null/v5"
	"testing"
	"time"
)

func TestSomeStructA_ToMap(t *testing.T) {
	amount := 123.45
	currency := "USD"
	createdAt := null.NewTime(time.Now(), true)
	updatedAt := null.NewTime(time.Now(), true)

	sa := SomeStructA{
		Id:                   1,
		Name:                 "Test Name",
		Age:                  30,
		TotalAmount:          1000.50,
		IsActive:             true,
		WhatTheFuckIsComplex: 3 + 4i,
		AmountCurrency: AmountCurrency{
			Amount:   12,
			Currency: "IDR",
		},
		Amount:    &amount,
		Currency:  &currency,
		CreatedBy: "Creator",
		UpdatedBy: "Updater",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	result := make(map[string]any, sa.FieldsLen())
	sa.ToMap(result)

	expected := map[string]interface{}{
		"id":                       uint(1),
		"name":                     "Test Name",
		"age":                      30,
		"total_amount":             1000.50,
		"is_active":                true,
		"what_the_fuck_is_complex": complex64(3 + 4i),
		"amount":                   amount,
		"amount_currency":          nil,
		"currency":                 currency,
		"created_by":               "Creator",
		"updated_by":               "Updater",
		"created_at":               createdAt.Time,
		"updated_at":               updatedAt.Time,
	}

	for key, value := range expected {
		if result[key] != value {
			t.Errorf("Expected %v for key %s but got %v", value, key, result[key])
		}
	}
}

func BenchmarkSomeStructA_ToMap(b *testing.B) {
	amount := 123.45
	currency := "USD"
	createdAt := null.NewTime(time.Now(), true)
	updatedAt := null.NewTime(time.Now(), true)

	sa := SomeStructA{
		Id:                   1,
		Name:                 "Test Name",
		Age:                  30,
		TotalAmount:          1000.50,
		IsActive:             true,
		WhatTheFuckIsComplex: 3 + 4i,
		Amount:               &amount,
		Currency:             &currency,
		CreatedBy:            "Creator",
		UpdatedBy:            "Updater",
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
	}

	result := make(map[string]any, sa.FieldsLen())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sa.ToMap(result)
	}
}

func BenchmarkSomeStructB_ToMap(b *testing.B) {
	amount := 123.45
	currency := "USD"
	sa := SomeStructB{
		Id:        1,
		Name:      "Test Name",
		Amount:    &amount,
		Currency:  &currency,
		CreatedBy: "Creator",
		UpdatedBy: "Updater",
		CreatedAt: null.TimeFrom(time.Now()),
		UpdatedAt: null.TimeFrom(time.Now()),
	}

	result := make(map[string]any, sa.FieldsLen())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sa.ToMap(result)
	}
}
