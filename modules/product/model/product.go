package model

import (
	"errors"
	"time"
)

const (
	EntityName = "product"
)

var (
	ErrorCreateProduct             = errors.New("create product failed")
	ErrorFindAllProduct            = errors.New("find all product failed")
	ErrorFindProductNotFound       = errors.New("find product not found")
	ErrorUpdateProductIsNotChanged = errors.New("update product is not changed")
)

// Create Product
type ProductCreationType struct {
	ProductNameVN            string    `json:"product_name_vn" validate:"required"`
	ProductNameENG           string    `json:"product_name_eng" validate:"required"`
	ProductNameDE            string    `json:"product_name_de" validate:"required"`
	ProductNameTH            string    `json:"product_name_th" validate:"required"`
	ProductCode              string    `json:"product_code" validate:"required"`
	Dph                      int32     `json:"dph"`
	Description              string    `json:"description"`
	ImageURL                 string    `json:"image_url"`
	Status                   string    `json:"status"`
	SubCategoryID            string    `json:"sub_category_id" validate:"required"`
	MinimumOrderQuantity     int       `json:"minimum_order_quantity"`
	MaximumOrderQuantity     int       `json:"maximum_order_quantity"`
	ReorderLevel             int       `json:"reorder_level"`
	IsStackability           bool      `json:"is_stackability"`
	TemperatureRequirement   float64   `json:"temperature_requirement"`
	IsFragility              bool      `json:"is_fragility"`
	ShelfLife                int       `json:"shelf_life" validate:"required"`
	Note                     string    `json:"note"`
	Season                   string    `json:"season"`
	IsPublished              bool      `json:"is_published"`
	PublishedAt              time.Time `json:"published_at"`
	PreOrder                 time.Time `json:"pre_order" validate:"required"`
	Length                   string    `json:"length"`
	Width                    float64   `json:"width"`
	Height                   float64   `json:"height"`
	NetWeight                float64   `json:"net_weight" validate:"required"`
	GrossWeight              float64   `json:"gross_weight" validate:"required"`
	Cubic                    float64   `json:"cubic" validate:"required"`
	AttitudeProductPackageID string    `json:"attitude_product_package_id"`
	BrandID                  string    `json:"brand_id"`
	TotalQuantity            int       `json:"total_quantity" validate:"required"`
}

type ProductType struct {
	ProductID                string    `json:"product_id" validate:"required"`
	ProductNameVN            string    `json:"product_name_vn" validate:"required"`
	ProductNameENG           string    `json:"product_name_eng" validate:"required"`
	ProductNameDE            string    `json:"product_name_de" validate:"required"`
	ProductNameTH            string    `json:"product_name_th" validate:"required"`
	ProductCode              string    `json:"product_code" validate:"required"`
	Dph                      int32     `json:"dph"`
	Description              string    `json:"description"`
	ImageURL                 string    `json:"image_url"`
	Status                   string    `json:"status"`
	SubCategoryID            string    `json:"sub_category_id" validate:"required"`
	MinimumOrderQuantity     int       `json:"minimum_order_quantity"`
	MaximumOrderQuantity     int       `json:"maximum_order_quantity"`
	ReorderLevel             int       `json:"reorder_level"`
	IsStackability           bool      `json:"is_stackability"`
	TemperatureRequirement   float64   `json:"temperature_requirement"`
	IsFragility              bool      `json:"is_fragility"`
	ShelfLife                int       `json:"shelf_life" validate:"required"`
	Note                     string    `json:"note"`
	Season                   string    `json:"season"`
	IsPublished              bool      `json:"is_published"`
	PublishedAt              time.Time `json:"published_at"`
	PreOrder                 time.Time `json:"pre_order" validate:"required"`
	Length                   string    `json:"length"`
	Width                    float64   `json:"width"`
	Height                   float64   `json:"height"`
	NetWeight                float64   `json:"net_weight" validate:"required"`
	GrossWeight              float64   `json:"gross_weight" validate:"required"`
	Cubic                    float64   `json:"cubic" validate:"required"`
	AttitudeProductPackageID string    `json:"attitude_product_package_id"`
	BrandID                  string    `json:"brand_id"`
	TotalQuantity            int       `json:"total_quantity" validate:"required"`
}
