package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/category/model"
)

type UpdateCategoryStorage interface {
	SelectCategoryById(cond map[string]interface{}) (*model.CategoryType, error)
	UpdateOldCategory(cond map[string]interface{}) error
}

type updateCategoryBiz struct {
	store UpdateCategoryStorage
}

func NewUpdateCategoryBiz(store UpdateCategoryStorage) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategory(newCategory *model.CategoryType) (*model.CategoryType, error) {
	//check old category
	oldCategory, err := biz.store.SelectCategoryById(map[string]interface{}{"category_id": newCategory.CategoryID})
	if err != nil {
		return nil, commonError.ErrCannotFindCategoryById(model.EntityName, err)
	}
	dataUpdate := getUpdatedFields(oldCategory, newCategory)
	//check dataUpdate is not changed
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrUpdateCategoryIsNotChanged(model.EntityName, model.ErrorUpdateCategoryIsNotChanged)
	}
	//update category
	if err := biz.store.UpdateOldCategory(dataUpdate); err != nil {
		return nil, commonError.ErrUpdateCategory(model.EntityName, err)
	}
	//get new category is updated
	newCategoryUpdated, err := biz.store.SelectCategoryById(map[string]interface{}{"category_id": newCategory.CategoryID})
	if err != nil {
		return nil, commonError.ErrCannotFindCategoryById(model.EntityName, err)
	}
	return newCategoryUpdated, nil
}

func getUpdatedFields(oldCategory *model.CategoryType, newCategory *model.CategoryType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldCategory.CategoryID == newCategory.CategoryID {
		updates["CategoryID"] = newCategory.CategoryID
	}
	if oldCategory.CategoryNameENG != newCategory.CategoryNameENG {
		updates["CategoryNameENG"] = newCategory.CategoryNameENG
	}
	if oldCategory.CategoryNameVN != newCategory.CategoryNameVN {
		updates["CategoryNameVN"] = newCategory.CategoryNameVN
	}
	if oldCategory.CategoryNameDE != newCategory.CategoryNameDE {
		updates["CategoryNameDE"] = newCategory.CategoryNameDE
	}
	if oldCategory.CategoryNameTH != newCategory.CategoryNameTH {
		updates["CategoryNameTH"] = newCategory.CategoryNameTH
	}
	if oldCategory.ImageURL != newCategory.ImageURL {
		updates["ImageURL"] = newCategory.ImageURL
	}
	return updates
}
