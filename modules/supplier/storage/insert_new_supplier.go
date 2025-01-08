package storage

import (
	"fmt"
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/supplier/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewSupplier(data *model.SupplierCreationType) error {
	id := uuid.NewString()
	in_put := schema.Supplier{
		SupplierID:    id,
		SupplierName:  data.SupplierName,
		SupplierCode:  data.SupplierCode,
		TaxID:         data.TaxID,
		Description:   data.Description,
		CurrencyID:    data.CurrencyID,
		EmailPurchase: data.EmailPurchase,
		Note:          data.Note,
		ContactInfo:   data.ContactInfo,
	}
	fmt.Println(in_put)
	tx := s.db.Begin()
	if err := tx.Create(&in_put).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
