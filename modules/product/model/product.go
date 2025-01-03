package model

import "errors"

const (
	EntityName = "product"
)

var (
	ErrorCreateProduct = errors.New("create product failed")
)

// Create Product
type ProductCreationType struct {
	ProductName string `json:"product_name"`
	Description string `json:"description"`
}
