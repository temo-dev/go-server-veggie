package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/purchase-price/model"
)

func (s *sqlStore) SelectAllPurchasePrice() ([]*model.PurchasePriceType, error) {
	var listPurchasePrice []*model.PurchasePriceType
	tx := s.db.Begin()
	result := tx.Table("purchase_prices").Scan(&listPurchasePrice)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listPurchasePrice, nil
}
