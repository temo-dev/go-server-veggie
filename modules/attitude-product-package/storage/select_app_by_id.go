package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/attitude-product-package/model"
)

func (s *sqlStore) SelectAPPById(cond map[string]interface{}) (*model.AttitudeProductPackageType, error) {
	var att *model.AttitudeProductPackageType
	tx := s.db.Begin()
	result := tx.Where("attitude_product_package_id = ?", cond["attitude_product_package_id"]).First(&schema.AttitudeProductPackage{})
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	result.Scan(&att)
	tx.Commit()
	return att, nil
}
