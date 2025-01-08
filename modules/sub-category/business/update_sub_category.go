package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/sub-category/model"
)

type UpdateSubCategoryStorage interface {
	SelectSubCategoryById(cond map[string]interface{}) (*model.SubCategoryType, error)
	UpdateOldSubCategory(cond map[string]interface{}) error
}

type updateSubCategoryBiz struct {
	store UpdateSubCategoryStorage
}

func NewUpdateSubCategoryBiz(store UpdateSubCategoryStorage) *updateSubCategoryBiz {
	return &updateSubCategoryBiz{store: store}
}

func (biz *updateSubCategoryBiz) UpdateSubCategory(newSubCategory *model.SubCategoryType) (*model.SubCategoryType, error) {
	//check old category
	oldSubCategory, err := biz.store.SelectSubCategoryById(map[string]interface{}{"sub_category_id": newSubCategory.SubCategoryID})
	if err != nil {
		return nil, commonError.ErrCannotFindSubCategoryById(model.EntityName, err)
	}
	dataUpdate := getUpdatedFields(oldSubCategory, newSubCategory)
	//check dataUpdate is not changed
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrUpdateSubCategoryIsNotChanged(model.EntityName, model.ErrorUpdateSubCategoryIsNotChanged)
	}
	//update category
	if err := biz.store.UpdateOldSubCategory(dataUpdate); err != nil {
		return nil, commonError.ErrUpdateSubCategory(model.EntityName, err)
	}
	//get new category is updated
	newSubCategoryUpdated, err := biz.store.SelectSubCategoryById(map[string]interface{}{"sub_category_id": newSubCategory.SubCategoryID})
	if err != nil {
		return nil, commonError.ErrCannotFindSubCategoryById(model.EntityName, err)
	}
	return newSubCategoryUpdated, nil
}

func getUpdatedFields(oldSubCategory *model.SubCategoryType, newSubCategory *model.SubCategoryType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldSubCategory.SubCategoryID == newSubCategory.SubCategoryID {
		updates["SubCategoryID"] = newSubCategory.SubCategoryID
	}
	if oldSubCategory.SubCategoryNameENG != newSubCategory.SubCategoryNameENG {
		updates["SubCategoryNameENG"] = newSubCategory.SubCategoryNameENG
	}
	if oldSubCategory.SubCategoryNameVN != newSubCategory.SubCategoryNameVN {
		updates["SubCategoryNameVN"] = newSubCategory.SubCategoryNameVN
	}
	if oldSubCategory.SubCategoryNameDE != newSubCategory.SubCategoryNameDE {
		updates["SubCategoryNameDE"] = newSubCategory.SubCategoryNameDE
	}
	if oldSubCategory.SubCategoryNameTH != newSubCategory.SubCategoryNameTH {
		updates["SubCategoryNameTH"] = newSubCategory.SubCategoryNameTH
	}
	if oldSubCategory.Dph != newSubCategory.Dph {
		updates["Dph"] = newSubCategory.Dph
	}
	if oldSubCategory.ImageURL != newSubCategory.ImageURL {
		updates["ImageURL"] = newSubCategory.ImageURL
	}
	if oldSubCategory.CategoryID != newSubCategory.CategoryID {
		updates["CategoryID"] = newSubCategory.CategoryID
	}
	return updates
}
