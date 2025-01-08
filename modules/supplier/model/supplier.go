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
	SupplierName  string `json:"supplier_name" validate:"required"`
	SupplierCode  string `json:"supplier_code" validate:"required"`
	TaxID         string `json:"tax_id" validate:"required"`
	Description   string `json:"description"`
	CurrencyID    string `json:"currency_id" validate:"required"`
	EmailPurchase string `json:"email_purchase" validate:"required"`
	Note          string `json:"note"`
	ContactInfo   string `json:"contact_info"`
}

type SupplierType struct {
	SupplierID         string  `json:"supplier_id"`
	SupplierName       string  `json:"supplier_name"`
	SupplierCode       string  `json:"supplier_code"`
	TaxID              string  `json:"tax_id"`
	Description        string  `json:"description"`
	CurrencyID         string  `json:"currency_id"`
	EmailPurchase      string  `json:"email_purchase" validate:"required"`
	Note               string  `json:"note"`
	OutstandingBalance float64 `json:"outstanding_balance"`
	Status             string  `json:"status"`
	DurationPakage     int     `json:"duration_pakage"`
	ContactInfo        string  `json:"contact_info"`
	Rate               float64 `json:"rate"`
}
