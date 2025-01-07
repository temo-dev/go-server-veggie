package model

import "errors"

const (
	EntityName = "supplier"
)

var (
	ErrorCreateSupplier             = errors.New("create supllier failed")
	ErrorSupplierNotFound           = errors.New("supplier not found")
	ErrorUpdateSupplierIsNotChanged = errors.New("update supplier is not changed")
)

// Create Supplier
type SupplierCreationType struct {
	SupplierName string `json:"supplier_name" validate:"required"`
	SupplierCode string `json:"supplier_code" validate:"required"`
	TaxID        string `json:"tax_id" validate:"required"`
	Description  string `json:"description"`
	CurrencyID   string `json:"currency_id" validate:"required"`
}

type SupplierType struct {
	SupplierID         string  `json:"supplier_id"`
	SupplierName       string  `json:"supplier_name"`
	SupplierCode       string  `json:"supplier_code"`
	TaxID              string  `json:"tax_id"`
	Description        string  `json:"description"`
	CurrencyID         string  `json:"currency_id"`
	OutstandingBalance float64 `json:"outstanding_balance"`
	Status             string  `json:"status"`
}
