package business

import (
	commonError "server-veggie/common/error"
	model "server-veggie/modules/category/model"
)

type DeleteCategoryByIdStorage interface {
	SelectCategoryById(cond map[string]interface{}) (*model.CategoryType, error)
	DeleteCategory(cond map[string]interface{}) error
}
type deleteCategoryByIdBiz struct {
	store DeleteCategoryByIdStorage
}

func NewDeleteCategoryByIdBiz(store DeleteCategoryByIdStorage) *deleteCategoryByIdBiz {
	return &deleteCategoryByIdBiz{store: store}
}

func (biz *deleteCategoryByIdBiz) DeleteCategoryById(id string) error {
	category, err := biz.store.SelectCategoryById(map[string]interface{}{"category_id": id})
	if err != nil {
		return err
	}
	if category == nil {
		return commonError.ErrCannotFindCategoryById(model.EntityName, model.ErrorFindCategoryById)
	}
	if err := biz.store.DeleteCategory(map[string]interface{}{"category_id": id}); err != nil {
		return err
	}
	return nil
}
