package business

import (
	"server-veggie/modules/sub-category/model"
)

type FindSubCategoryByIdStorage interface {
	SelectSubCategoryById(cond map[string]interface{}) (*model.SubCategoryType, error)
}

type findSubCategoryByIdBiz struct {
	store FindSubCategoryByIdStorage
}

func NewFindSubCategoryByIdBiz(store FindSubCategoryByIdStorage) *findSubCategoryByIdBiz {
	return &findSubCategoryByIdBiz{store: store}
}

func (biz *findSubCategoryByIdBiz) FindSubCategoryById(id string) (*model.SubCategoryType, error) {
	category, err := biz.store.SelectSubCategoryById(map[string]interface{}{"sub_category_id": id})
	if err != nil {
		return nil, err
	}
	return category, nil
}
