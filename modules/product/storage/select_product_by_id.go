package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/product/model"
)

func (s *sqlStore) SelectProductById(cond map[string]interface{}) (*model.ProductType, error) {
	var product *model.ProductType
	tx := s.db.Begin()
	result := tx.Where("product_id = ?", cond["product_id"]).First(&schema.Product{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&product)
	tx.Commit()
	return product, nil
}
