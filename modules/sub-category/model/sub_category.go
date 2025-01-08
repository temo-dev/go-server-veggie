package model

import "errors"

const (
	EntityName = "SubCategory"
)

var (
	ErrorCreateSubCategory             = errors.New("create subcategory failed")
	ErrorUpdateSubCategoryIsNotChanged = errors.New("update subcategory is not changed")
	ErrorFindSubCategoryById           = errors.New("find subcategory is not exist")
)

// Create Category
type SubCategoryCreationType struct {
	SubCategoryNameVN  string `json:"sub_category_name_vn" validate:"required"`
	SubCategoryNameENG string `json:"sub_category_name_eng" validate:"required"`
	SubCategoryNameDE  string `json:"sub_category_name_de" validate:"required"`
	SubCategoryNameTH  string `json:"sub_category_name_th" validate:"required"`
	CategoryID         string `json:"category_id" validate:"required"`
	ImageURL           string `json:"image_url"`
	Dph                int32  `json:"dph"`
}

// Get Categorie
type SubCategoryType struct {
	SubCategoryID      string `json:"sub_category_id" validate:"required"`
	SubCategoryNameVN  string `json:"sub_category_name_vn" validate:"required"`
	SubCategoryNameENG string `json:"sub_category_name_eng" validate:"required"`
	SubCategoryNameDE  string `json:"sub_category_name_de" validate:"required"`
	SubCategoryNameTH  string `json:"sub_category_name_th" validate:"required"`
	CategoryID         string `json:"category_id" validate:"required"`
	ImageURL           string `json:"image_url"`
	Dph                int32  `json:"dph"`
}
