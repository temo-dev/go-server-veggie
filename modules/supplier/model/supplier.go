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
}
