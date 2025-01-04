package model

import "errors"

const (
	EntityName = "supplier"
)

var (
	ErrorCreateSupplier = errors.New("create supllier failed")
)

// Create Supplier
type SupplierCreationType struct {
	SupplierName string `json:"supplier_name" validate:"required"`
	TaxID        string `json:"tax_id" validate:"required"`
	Description  string `json:"description"`
	CurrencyID   string `gorm:"type:uuid;not null"`
}
