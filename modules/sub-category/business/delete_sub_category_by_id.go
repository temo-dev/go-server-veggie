package business

import (
	commonError "server-veggie/common/error"
	model "server-veggie/modules/sub-category/model"
)

type DeleteSubCategoryByIdStorage interface {
	SelectSubCategoryById(cond map[string]interface{}) (*model.SubCategoryType, error)
	DeleteSubCategory(cond map[string]interface{}) error
}
type deleteSubCategoryByIdBiz struct {
	store DeleteSubCategoryByIdStorage
}

func NewDeleteSubCategoryByIdBiz(store DeleteSubCategoryByIdStorage) *deleteSubCategoryByIdBiz {
	return &deleteSubCategoryByIdBiz{store: store}
}

func (biz *deleteSubCategoryByIdBiz) DeleteSubCategoryById(id string) error {
	category, err := biz.store.SelectSubCategoryById(map[string]interface{}{"sub_category_id": id})
	if err != nil {
		return err
	}
	if category == nil {
		return commonError.ErrCannotFindSubCategoryById(model.EntityName, model.ErrorFindSubCategoryById)
	}
	if err := biz.store.DeleteSubCategory(map[string]interface{}{"sub_category_id": id}); err != nil {
		return err
	}
	return nil
}
