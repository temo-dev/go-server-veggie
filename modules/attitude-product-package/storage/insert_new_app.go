package storage

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/attitude-product-package/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertNewAPP(data *model.AttitudeProductPackageCreationType) error {
	id := uuid.NewString()
	att := schema.AttitudeProductPackage{
		AttitudeProductPackageID:   id,
		AttitudeProductPackageCode: data.AttitudeProductPackageCode,
		PackageLength:              data.PackageLength,
		PackageWidth:               data.PackageWidth,
		PackageHeight:              data.PackageHeight,
		PackageCubic:               data.PackageCubic,
	}
	tx := s.db.Begin()
	if err := tx.Create(&att).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
