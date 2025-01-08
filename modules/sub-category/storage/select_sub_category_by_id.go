package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/sub-category/model"
)

func (s *sqlStore) SelectSubCategoryById(cond map[string]interface{}) (*model.SubCategoryType, error) {
	var subcategory *model.SubCategoryType
	tx := s.db.Begin()
	result := tx.Where("sub_category_id = ?", cond["sub_category_id"]).First(&schema.SubCategory{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&subcategory)
	tx.Commit()
	return subcategory, nil
}
