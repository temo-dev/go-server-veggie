package business

import (
	commonError "server-veggie/common/error"
	model "server-veggie/modules/category/model"
)

type FindCategoriesStorage interface {
	SelectCategories() ([]*model.CategoryType, error)
}

type findCategoriesBiz struct {
	store FindCategoriesStorage
}

func NewFindCategoriesBiz(store FindCategoriesStorage) *findCategoriesBiz {
	return &findCategoriesBiz{store: store}
}

func (biz *findCategoriesBiz) FindCategories() ([]*model.CategoryType, error) {
	//hanlde business
	listCategories, err := biz.store.SelectCategories()
	if err != nil {
		return nil, commonError.ErrCannotFindCategories(model.EntityName, err)
	}
	return listCategories, nil
}
