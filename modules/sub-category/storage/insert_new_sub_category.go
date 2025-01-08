package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/sub-category/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewSubCategory(data *model.SubCategoryCreationType) error {
	id := uuid.NewString()
	category := schema.SubCategory{
		SubCategoryID:      id,
		SubCategoryNameVN:  data.SubCategoryNameVN,
		SubCategoryNameENG: data.SubCategoryNameENG,
		SubCategoryNameDE:  data.SubCategoryNameDE,
		SubCategoryNameTH:  data.SubCategoryNameTH,
		CategoryID:         data.CategoryID,
		ImageURL:           data.ImageURL,
		Dph:                data.Dph,
	}
	tx := s.db.Begin()
	if err := tx.Create(&category).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
