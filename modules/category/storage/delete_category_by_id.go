package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteCategory(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("category_id", cond["category_id"]).Delete(&schema.Category{}).Error; err != nil {
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
