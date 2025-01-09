package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
)

func (s *sqlStore) UpdateOldAPP(cond map[string]interface{}) error {
	tx := s.db.Begin()
	if err := tx.Model(&schema.AttitudeProductPackage{}).Where("attitude_product_package_id", cond["attitude_product_package_id"]).Updates(cond).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	return nil
}
