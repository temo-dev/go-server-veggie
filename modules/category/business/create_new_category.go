package business

import (
	commonError "server-veggie/common/error"
	model "server-veggie/modules/category/model"
)

type CreateCategoryStorage interface {
	InsertNewCategory(data *model.CategoryCreationType) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateNewCategory(data *model.CategoryCreationType) error {
	if err := biz.store.InsertNewCategory(data); err != nil {
		return commonError.ErrCannotCreateCategory(model.EntityName, err)
	}
	return nil
}
