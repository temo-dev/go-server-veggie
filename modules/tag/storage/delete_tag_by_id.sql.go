package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteTagById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("tag_id", cond["tag_id"]).Delete(&schema.Tag{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()

	return nil
}
