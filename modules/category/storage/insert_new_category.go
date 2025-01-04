package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/category/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewCategory(data *model.CategoryCreationType) error {
	id := uuid.NewString()
	category := schema.Category{
		CategoryID:      id,
		CategoryNameVN:  data.CategoryNameVN,
		CategoryNameENG: data.CategoryNameENG,
		Dph:             data.Dph,
		ImageURL:        data.ImageURL,
	}
	tx := s.db.Begin()
	if err := tx.Create(&category).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
