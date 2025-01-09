package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/attitude-product-package/model"
)

type UpdateOldAPPStorage interface {
	SelectAPPById(cond map[string]interface{}) (*model.AttitudeProductPackageType, error)
	UpdateOldAPP(cond map[string]interface{}) error
}

type updateOldAPPBiz struct {
	store UpdateOldAPPStorage
}

func NewUpdateOldAPPBiz(store UpdateOldAPPStorage) *updateOldAPPBiz {
	return &updateOldAPPBiz{store: store}
}

func (biz *updateOldAPPBiz) UpdateOldAPP(newAPP *model.AttitudeProductPackageType) (*model.AttitudeProductPackageType, error) {
	//check old app
	oldAPP, err := biz.store.SelectAPPById(map[string]interface{}{"attitude_product_package_id": newAPP.AttitudeProductPackageID})
	if err != nil {
		return nil, err
	}
	dataUpdate := getUpdatedFields(oldAPP, newAPP)
	//check dataUpdate is not changed
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrCannotUpdateAPP(model.EntityName, model.ErrorUpdateAPPIsNotChanged)
	}
	//update app
	if err := biz.store.UpdateOldAPP(dataUpdate); err != nil {
		return nil, err
	}
	//get new app is updated
	newAPPUpdated, err := biz.store.SelectAPPById(map[string]interface{}{"attitude_product_package_id": newAPP.AttitudeProductPackageID})
	if err != nil {
		return nil, err
	}
	return newAPPUpdated, nil
}

func getUpdatedFields(oldAPP *model.AttitudeProductPackageType, newAPP *model.AttitudeProductPackageType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldAPP.AttitudeProductPackageID == newAPP.AttitudeProductPackageID {
		updates["AttitudeProductPackageID"] = newAPP.AttitudeProductPackageID
	}
	if oldAPP.AttitudeProductPackageCode != newAPP.AttitudeProductPackageCode {
		updates["AttitudeProductPackageCode"] = newAPP.AttitudeProductPackageCode
	}
	if oldAPP.PackageCubic != newAPP.PackageCubic {
		updates["PackageCubic"] = newAPP.PackageCubic
	}
	if oldAPP.PackageGrossWeight != newAPP.PackageGrossWeight {
		updates["PackageGrossWeight"] = newAPP.PackageGrossWeight
	}
	if oldAPP.PackageHeight != newAPP.PackageHeight {
		updates["PackageHeight"] = newAPP.PackageHeight
	}
	if oldAPP.PackageNetWeight != newAPP.PackageNetWeight {
		updates["PackageNetWeight"] = newAPP.PackageNetWeight
	}
	if oldAPP.PackageLength != newAPP.PackageLength {
		updates["PackageLength"] = newAPP.PackageLength
	}
	if oldAPP.PackageWidth != newAPP.PackageWidth {
		updates["PackageWidth"] = newAPP.PackageWidth
	}
	if oldAPP.UnitsPerBox != newAPP.UnitsPerBox {
		updates["UnitsPerBox"] = newAPP.UnitsPerBox
	}
	return updates
}
