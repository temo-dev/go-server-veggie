package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/category/model"
)

func (s *sqlStore) SelectCategories() ([]*model.CategoryType, error) {
	var listCategories []*model.CategoryType
	tx := s.db.Begin()
	result := tx.Table("categories").Scan(&listCategories)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listCategories, nil
}
