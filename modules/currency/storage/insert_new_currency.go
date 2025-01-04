package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/currency/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewCurrency(data *model.CurrencyCreationType) error {
	id := uuid.NewString()
	in_put := schema.Currency{
		CurrencyID:   id,
		CurrencyName: data.CurrencyName,
		CurrencyCode: data.CurrencyCode,
		ExchangeRate: data.ExchangeRate,
	}
	tx := s.db.Begin()
	if err := tx.Create(&in_put).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
