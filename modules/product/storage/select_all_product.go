package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/product/model"
)

func (s *sqlStore) SelectAllProduct() ([]*model.ProductType, error) {
	var listProducts []*model.ProductType
	tx := s.db.Begin()
	result := tx.Table("products").Scan(&listProducts)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listProducts, nil
}
