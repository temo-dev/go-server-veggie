package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/category/model"
)

type DeleteCategoryByIdStorage interface {
	DeleteCategory(cond map[string]interface{}) error
}
type deleteCategoryByIdBiz struct {
	store DeleteCategoryByIdStorage
}

func NewDeleteCategoryByIdBiz(store DeleteCategoryByIdStorage) *deleteCategoryByIdBiz {
	return &deleteCategoryByIdBiz{store: store}
}

func (biz *deleteCategoryByIdBiz) DeleteCategoryById(id string) error {
	if err := biz.store.DeleteCategory(map[string]interface{}{"category_id": id}); err != nil {
		return commonError.ErrCannotDeleteCategoryById(model.EntityName, err)
	}
	return nil
}
