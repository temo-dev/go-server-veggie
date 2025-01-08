package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldSubCategory(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.SubCategory{}).Where("sub_category_id", cond["SubCategoryID"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
