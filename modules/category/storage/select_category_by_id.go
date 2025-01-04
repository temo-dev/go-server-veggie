package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/category/model"
)

func (s *sqlStore) SelectCategoryById(cond map[string]interface{}) (*model.CategoryType, error) {
	var category *model.CategoryType
	tx := s.db.Begin()
	result := tx.Where("category_id = ?", cond["category_id"]).First(&schema.Category{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&category)
	tx.Commit()
	return category, nil
}
