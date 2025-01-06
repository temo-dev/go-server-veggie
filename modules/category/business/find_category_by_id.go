package business

import (
	"server-veggie/modules/category/model"
)

type FindCategoryByIdStorage interface {
	SelectCategoryById(cond map[string]interface{}) (*model.CategoryType, error)
}

type findCategoryByIdBiz struct {
	store FindCategoryByIdStorage
}

func NewFindCategoryByIdBiz(store FindCategoryByIdStorage) *findCategoryByIdBiz {
	return &findCategoryByIdBiz{store: store}
}

func (biz *findCategoryByIdBiz) FindCategoryById(id string) (*model.CategoryType, error) {
	category, err := biz.store.SelectCategoryById(map[string]interface{}{"category_id": id})
	if err != nil {
		return nil, err
	}
	return category, nil
}
