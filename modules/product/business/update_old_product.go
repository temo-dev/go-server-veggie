package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/product/model"
)

type UpdateOldProductStorage interface {
	SelectProductById(cond map[string]interface{}) (*model.ProductType, error)
	UpdateOldProduct(cond map[string]interface{}) error
}
type updateOldProductBiz struct {
	store UpdateOldProductStorage
}

func NewUpdateOldProductBiz(store UpdateOldProductStorage) *updateOldProductBiz {
	return &updateOldProductBiz{store: store}
}

func (biz *updateOldProductBiz) UpdateOldProduct(newProduct *model.ProductType) (*model.ProductType, error) {
	//check old product
	oldProduct, err := biz.store.SelectProductById(map[string]interface{}{"product_id": newProduct.ProductID})
	if err != nil {
		return nil, commonError.ErrCannotUpdateProduct(model.EntityName, err)
	}
	dataUpdate := getUpdatedFields(oldProduct, newProduct)
	//check dataUpdate is not changed
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrCannotUpdateProduct(model.EntityName, model.ErrorUpdateProductIsNotChanged)
	}
	//update product
	if err := biz.store.UpdateOldProduct(dataUpdate); err != nil {
		return nil, commonError.ErrCannotUpdateProduct(model.EntityName, err)
	}
	//get new product is updated
	newProductUpdated, err := biz.store.SelectProductById(map[string]interface{}{"product_id": newProduct.ProductID})
	if err != nil {
		return nil, commonError.ErrCannotUpdateProduct(model.EntityName, err)
	}
	return newProductUpdated, nil
}

func getUpdatedFields(oldProduct *model.ProductType, newProduct *model.ProductType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldProduct.ProductID == newProduct.ProductID {
		updates["ProductID"] = newProduct.ProductID
	}
	if oldProduct.ProductNameVN != newProduct.ProductNameVN {
		updates["ProductNameVN"] = newProduct.ProductNameVN
	}
	if oldProduct.ProductNameENG != newProduct.ProductNameENG {
		updates["ProductNameENG"] = newProduct.ProductNameENG
	}
	if oldProduct.ProductNameDE != newProduct.ProductNameDE {
		updates["ProductNameDE"] = newProduct.ProductNameDE
	}
	if oldProduct.ProductNameTH != newProduct.ProductNameTH {
		updates["ProductNameTH"] = newProduct.ProductNameTH
	}
	if oldProduct.ProductCode != newProduct.ProductCode {
		updates["ProductCode"] = newProduct.ProductCode
	}
	if oldProduct.Dph != newProduct.Dph {
		updates["Dph"] = newProduct.Dph
	}
	if oldProduct.Description != newProduct.Description {
		updates["Description"] = newProduct.Description
	}
	if oldProduct.ImageURL != newProduct.ImageURL {
		updates["ImageURL"] = newProduct.ImageURL
	}
	if oldProduct.Status != newProduct.Status {
		updates["Status"] = newProduct.Status
	}
	if oldProduct.SubCategoryID != newProduct.SubCategoryID {
		updates["SubCategoryID"] = newProduct.SubCategoryID
	}
	if oldProduct.MinimumOrderQuantity != newProduct.MinimumOrderQuantity {
		updates["MinimumOrderQuantity"] = newProduct.MinimumOrderQuantity
	}
	if oldProduct.MaximumOrderQuantity != newProduct.MaximumOrderQuantity {
		updates["MaximumOrderQuantity"] = newProduct.MaximumOrderQuantity
	}
	if oldProduct.ReorderLevel != newProduct.ReorderLevel {
		updates["ReorderLevel"] = newProduct.ReorderLevel
	}
	if oldProduct.IsStackability != newProduct.IsStackability {
		updates["IsStackability"] = newProduct.IsStackability
	}
	if oldProduct.TemperatureRequirement != newProduct.TemperatureRequirement {
		updates["TemperatureRequirement"] = newProduct.TemperatureRequirement
	}
	if oldProduct.TemperatureRequirement != newProduct.TemperatureRequirement {
		updates["TemperatureRequirement"] = newProduct.TemperatureRequirement
	}
	if oldProduct.IsFragility != newProduct.IsFragility {
		updates["IsFragility"] = newProduct.IsFragility
	}
	if oldProduct.ShelfLife != newProduct.ShelfLife {
		updates["ShelfLife"] = newProduct.ShelfLife
	}
	if oldProduct.Note != newProduct.Note {
		updates["Note"] = newProduct.Note
	}
	if oldProduct.Season != newProduct.Season {
		updates["Season"] = newProduct.Season
	}
	if oldProduct.IsPublished != newProduct.IsPublished {
		updates["IsPublished"] = newProduct.IsPublished
	}
	if oldProduct.PublishedAt != newProduct.PublishedAt {
		updates["ShelfLife"] = newProduct.ShelfLife
	}
	if oldProduct.PreOrder != newProduct.PreOrder {
		updates["ShelfLife"] = newProduct.ShelfLife
	}
	if oldProduct.Length != newProduct.Length {
		updates["Length"] = newProduct.Length
	}
	if oldProduct.Width != newProduct.Width {
		updates["Width"] = newProduct.Width
	}
	if oldProduct.Height != newProduct.Height {
		updates["Height"] = newProduct.Height
	}
	if oldProduct.NetWeight != newProduct.NetWeight {
		updates["ShelfLife"] = newProduct.ShelfLife
	}
	if oldProduct.GrossWeight != newProduct.GrossWeight {
		updates["GrossWeight"] = newProduct.GrossWeight
	}
	if oldProduct.Cubic != newProduct.Cubic {
		updates["Cubic"] = newProduct.Cubic
	}
	if oldProduct.BrandID != newProduct.BrandID {
		updates["BrandID"] = newProduct.BrandID
	}
	if oldProduct.AttitudeProductPackageID != newProduct.AttitudeProductPackageID {
		updates["AttitudeProductPackageID"] = newProduct.AttitudeProductPackageID
	}
	if oldProduct.TotalQuantity != newProduct.TotalQuantity {
		updates["TotalQuantity"] = newProduct.TotalQuantity
	}
	return updates
}
