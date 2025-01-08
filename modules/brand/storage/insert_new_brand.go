package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/brand/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewBrand(brand *model.BrandCreationtype) error {
	id := uuid.NewString()
	in_put := schema.Brand{
		BrandID:     id,
		BrandName:   brand.BrandName,
		Description: brand.Description,
	}
	tx := s.db.Begin()
	if err := tx.Create(&in_put).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
