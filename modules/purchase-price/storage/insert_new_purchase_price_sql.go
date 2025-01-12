package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/purchase-price/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewPurchasePrice(data *model.PurchasePriceCreationType) error {
	id := uuid.NewString()
	purchasePrice := schema.PurchasePrice{
		PurchasePriceID: id,
		ProductID:       data.ProductID,
		CurrencyID:      data.CurrencyID,
		SupplierID:      data.SupplierID,
		Season:          data.Season,
		RetailPrice:     data.RetailPrice,
		BoxPrice:        data.BoxPrice,
		PalletPrice:     data.PalletPrice,
		ContainerPrice:  data.ContainerPrice,
	}
	tx := s.db.Begin()
	if err := tx.Create(&purchasePrice).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
