package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/attitude-product-package/model"
)

func (s *sqlStore) SelectAllAPP() ([]*model.AttitudeProductPackageType, error) {
	var listApp []*model.AttitudeProductPackageType
	tx := s.db.Begin()
	result := tx.Table("attitude_product_packages").Scan(&listApp)
	if result.Error != nil {
		tx.Rollback()
		return nil, commonError.ErrDB(result.Error)
	}
	tx.Commit()
	return listApp, nil
}
