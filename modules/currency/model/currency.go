package model

import "errors"

const (
	EntityName = "currency"
)

var (
	ErrorCreateCurrency   = errors.New("create currency failed")
	ErrorFindCurrencyById = errors.New("find currency failed by id")
	ErrorDeleteCurrency   = errors.New("delete currency failed")
)

// Create Currency
type CurrencyCreationType struct {
	CurrencyName string  `json:"currency_name" validate:"required"`
	CurrencyCode string  `json:"currency_code" validate:"required"`
	ExchangeRate float64 `json:"exchange_rate"`
}

type CurencyType struct {
	CurrencyID   string  `json:"currency_id"`
	CurrencyName string  `json:"currency_name"`
	CurrencyCode string  `json:"currency_code"`
	ExchangeRate float64 `json:"exchange_rate"`
}
