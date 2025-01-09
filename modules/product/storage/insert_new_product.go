package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/product/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewProduct(data *model.ProductCreationType) error {
	id := uuid.NewString()
	product := schema.Product{
		ProductID:                id,
		ProductNameVN:            data.ProductNameVN,
		ProductNameENG:           data.ProductNameENG,
		ProductNameDE:            data.ProductNameDE,
		ProductNameTH:            data.ProductNameTH,
		ProductCode:              data.ProductCode,
		Dph:                      data.Dph,
		Description:              data.Description,
		ImageURL:                 data.ImageURL,
		Status:                   data.Status,
		SubCategoryID:            data.SubCategoryID,
		MinimumOrderQuantity:     data.MinimumOrderQuantity,
		MaximumOrderQuantity:     data.MaximumOrderQuantity,
		ReorderLevel:             data.ReorderLevel,
		IsStackability:           data.IsStackability,
		TemperatureRequirement:   data.TemperatureRequirement,
		IsFragility:              data.IsFragility,
		ShelfLife:                data.ShelfLife,
		Note:                     data.Note,
		Season:                   data.Season,
		IsPublished:              data.IsPublished,
		PublishedAt:              data.PublishedAt,
		PreOrder:                 data.PreOrder,
		Length:                   data.Length,
		Width:                    data.Width,
		Height:                   data.Height,
		NetWeight:                data.NetWeight,
		GrossWeight:              data.GrossWeight,
		Cubic:                    data.Cubic,
		AttitudeProductPackageID: data.AttitudeProductPackageID,
		BrandID:                  data.BrandID,
		TotalQuantity:            data.TotalQuantity,
	}
	tx := s.db.Begin()
	if err := tx.Create(&product).Error; err != nil {
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
