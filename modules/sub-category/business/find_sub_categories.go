package business

import (
	model "server-veggie/modules/sub-category/model"
)

type FindSubCategoriesStorage interface {
	SelectSubCategories() ([]*model.SubCategoryType, error)
}

type findSubCategoriesBiz struct {
	store FindSubCategoriesStorage
}

func NewFindSubCategoriesBiz(store FindSubCategoriesStorage) *findSubCategoriesBiz {
	return &findSubCategoriesBiz{store: store}
}

func (biz *findSubCategoriesBiz) FindSubCategories() ([]*model.SubCategoryType, error) {
	//hanlde business
	listSubCategories, err := biz.store.SelectSubCategories()
	if err != nil {
		return nil, err
	}
	return listSubCategories, nil
}
