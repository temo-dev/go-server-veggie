package business

import (
	commonError "server-veggie/common/error"
	model "server-veggie/modules/sub-category/model"
)

type CreateSubCategoryStorage interface {
	InsertNewSubCategory(data *model.SubCategoryCreationType) error
}

type createSubCategoryBiz struct {
	store CreateSubCategoryStorage
}

func NewCreateSubCategoryBiz(store CreateSubCategoryStorage) *createSubCategoryBiz {
	return &createSubCategoryBiz{store: store}
}

func (biz *createSubCategoryBiz) CreateNewSubCategory(data *model.SubCategoryCreationType) error {
	if err := biz.store.InsertNewSubCategory(data); err != nil {
		return commonError.ErrCannotCreateCategory(model.EntityName, err)
	}
	return nil
}
