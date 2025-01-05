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
	ProductNameVN  string `json:"product_name_vn" validate:"required"`
	ProductNameENG string `json:"product_name_eng" validate:"required"`
	Description    string `json:"description"`
	ProductCode    string `json:"product_code" validate:"required"`
	Dph            int32  `json:"dph"`
	ImageURL       string `json:"image_url"`
	CategoryID     string `json:"category_id" validate:"required"`
}
