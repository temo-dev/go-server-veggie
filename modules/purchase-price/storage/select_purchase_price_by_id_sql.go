package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/purchase-price/model"
)

func (s *sqlStore) SelectPurchasePriceById(cond map[string]interface{}) (*model.PurchasePriceType, error) {
	var purchasePrice *model.PurchasePriceType
	tx := s.db.Begin()
	result := tx.Where("purchase_price_id = ?", cond["purchase_price_id"]).First(&schema.PurchasePrice{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&purchasePrice)
	tx.Commit()
	return purchasePrice, nil
}
