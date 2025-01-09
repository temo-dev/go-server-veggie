package model

import "errors"

const (
	EntityName = "AttitudeProductPackage"
)

var (
	ErrorCreateAPP             = errors.New("create attitude product package failed")
	ErrorUpdateAPPIsNotChanged = errors.New("update attitude product package is not changed")
	ErrorFindAPPById           = errors.New("find attitude product package is not exist")
)

// Create Category
type AttitudeProductPackageCreationType struct {
	AttitudeProductPackageCode string  `json:"attitude_product_package_code"`
	PackageLength              string  `json:"package_length" validate:"required"`
	PackageWidth               float64 `json:"package_width"`
	PackageHeight              float64 `json:"package_height"`
	PackageCubic               float64 `json:"package_cubic" validate:"required"`
}

// Get Categorie
type AttitudeProductPackageType struct {
	AttitudeProductPackageID   string  `json:"attitude_product_package_id" validate:"required"`
	AttitudeProductPackageCode string  `json:"attitude_product_package_code"`
	PackageLength              string  `json:"package_length" validate:"required"`
	PackageWidth               float64 `json:"package_width"`
	PackageHeight              float64 `json:"package_height"`
	PackageCubic               float64 `json:"package_cubic" validate:"required"`
}
