package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/sub-category/model"
)

func (s *sqlStore) SelectSubCategories() ([]*model.SubCategoryType, error) {
	var listSubCategories []*model.SubCategoryType
	tx := s.db.Begin()
	result := tx.Table("sub_categories").Scan(&listSubCategories)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listSubCategories, nil
}
