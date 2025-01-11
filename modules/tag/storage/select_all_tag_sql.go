package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/tag/model"
)

func (s *sqlStore) SelectAllTags() ([]*model.TagType, error) {
	var tags []*model.TagType
	tx := s.db.Begin()
	result := tx.Table("tags").Scan(&tags)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return tags, nil
}
