package model

import "errors"

const (
	EntityName = "category"
)

var (
	ErrorCreateCategory             = errors.New("create category failed")
	ErrorUpdateCategoryIsNotChanged = errors.New("update category is not changed")
	ErrorFindCategoryById           = errors.New("find category is not exist")
)

// Create Category
type CategoryCreationType struct {
	CategoryNameVN  string `json:"category_name_vn" validate:"required"`
	CategoryNameENG string `json:"category_name_eng" validate:"required"`
	CategoryNameDE  string `json:"category_name_de" validate:"required"`
	CategoryNameTH  string `json:"category_name_th" validate:"required"`
	ImageURL        string `json:"image_url"`
}

// Get Categorie
type CategoryType struct {
	CategoryID      string `json:"category_id" validate:"required"`
	CategoryNameVN  string `json:"category_name_vn" validate:"required"`
	CategoryNameENG string `json:"category_name_eng" validate:"required"`
	CategoryNameDE  string `json:"category_name_de" validate:"required"`
	CategoryNameTH  string `json:"category_name_th" validate:"required"`
	ImageURL        string `json:"image_url"`
}
