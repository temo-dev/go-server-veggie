package model

import (
	"errors"
	"time"
)

const (
	EntityName = "purchase_price"
)

var (
	ErrorPurchasePriceNotFound           = errors.New(" purchase price is not found")
	ErrorUpdatePurchasePriceIsNotChanged = errors.New("update purchase price is not changed")
)

// Create Purchase Price
type PurchasePriceCreationType struct {
	ProductID      string  `json:"product_id" validate:"required"`
	CurrencyID     string  `json:"currency_id" validate:"required"`
	SupplierID     string  `json:"supplier_id" validate:"required"`
	Season         string  `json:"season" validate:"required"`
	RetailPrice    float64 `json:"retail_price" validate:"required"`
	BoxPrice       float64 `json:"box_price" validate:"required"`
	PalletPrice    float64 `json:"pallet_price"`
	ContainerPrice float64 `json:"container_price"`
}

type PurchasePriceType struct {
	PurchasePriceID string    `json:"purchase_price_id" validate:"required"`
	ProductID       string    `json:"product_id" validate:"required"`
	CurrencyID      string    `json:"currency_id" validate:"required"`
	SupplierID      string    `json:"supplier_id" validate:"required"`
	Season          string    `json:"season" validate:"required"`
	RetailPrice     float64   `json:"retail_price" validate:"required"`
	BoxPrice        float64   `json:"box_price" validate:"required"`
	PalletPrice     float64   `json:"pallet_price" validate:"required"`
	ContainerPrice  float64   `json:"container_price" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
