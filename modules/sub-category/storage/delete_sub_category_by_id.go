package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteSubCategory(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("sub_category_id", cond["sub_category_id"]).Delete(&schema.SubCategory{}).Error; err != nil {
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
