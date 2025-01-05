package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/currency/model"
)

func (s *sqlStore) SelectAllCurrencies() (*[]model.CurencyType, error) {
	var listCurrencies *[]model.CurencyType
	tx := s.db.Begin()
	result := tx.Table("currencies").Scan(&listCurrencies)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listCurrencies, nil
}
