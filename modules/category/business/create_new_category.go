package business

import (
	"fmt"
	model "server-veggie/modules/category/model"
)

type CreateCategoryStorage interface {
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateNewCategory(data *model.CategoryCreationType) error {
	fmt.Println("data", data)
	return nil
}
