package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/tag/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewTag(data *model.TagCreationType) error {
	id := uuid.NewString()
	tag := schema.Tag{
		TagID:       id,
		TagName:     data.TagName,
		Description: data.Description,
		ImageURL:    data.ImageURL,
	}
	tx := s.db.Begin()
	if err := tx.Create(&tag).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
