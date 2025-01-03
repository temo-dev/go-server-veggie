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
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}
