package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/product/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewProduct(data *model.ProductCreationType) error {
	id := uuid.NewString()
	product := schema.Product{
		ProductID: id,
	}
	tx := s.db.Begin()
	if err := tx.Create(&product).Error; err != nil {
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
