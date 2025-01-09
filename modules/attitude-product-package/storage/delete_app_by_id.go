package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) DeleteAppById(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Where("attitude_product_package_id", cond["attitude_product_package_id"]).Delete(&schema.AttitudeProductPackage{}).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
