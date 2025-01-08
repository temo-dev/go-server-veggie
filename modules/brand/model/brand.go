package model

import "errors"

// error
const (
	EntityName = "brand"
)

var (
	ErrorCanotCreateBrand        = errors.New("cannot create brand")
	ErrorCanotUpdateBrand        = errors.New("cannot update brand")
	ErrorUpdateBrandIsNotChanged = errors.New("cannot update brand is not changed")
	ErrBrandNotFound             = errors.New("brand not found")
)

// model
type BrandCreationtype struct {
	BrandName   string `json:"brand_name" validate:"required"`
	Description string `json:"description"`
}

type BrandType struct {
	BrandID     string `json:"brand_id"`
	BrandName   string `json:"brand_name"`
	Description string `json:"description"`
}
