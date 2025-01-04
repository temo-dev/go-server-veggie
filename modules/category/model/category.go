package model

import "errors"

const (
	EntityName = "category"
)

var (
	ErrorCreateCategory = errors.New("create category failed")
)

// Create Category
type CategoryCreationType struct {
	CategoryNameVN  string `json:"category_name_vn" validate:"required"`
	CategoryNameENG string `json:"category_name_eng" validate:"required"`
	Dph             int32  `json:"dph" validate:"required"`
	ImageURL        string `json:"image_url"`
}

// Get Categorie
type CategoryType struct {
	CategoryID      string `json:"category_id"`
	CategoryNameVN  string `json:"category_name_vn"`
	CategoryNameENG string `json:"category_name_eng"`
	ImageURL        string `json:"image_url"`
	Dph             int32  `json:"dph"`
}
