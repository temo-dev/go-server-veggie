package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/tag/model"
)

func (s *sqlStore) SelectTagById(cond map[string]interface{}) (*model.TagType, error) {
	var tag *model.TagType
	tx := s.db.Begin()
	result := tx.Where("tag_id = ?", cond["tag_id"]).First(&schema.Tag{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&tag)
	tx.Commit()
	return tag, nil
}
