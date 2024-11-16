package business

import (
	"server-veggie/common"
	"server-veggie/modules/user/model"
)

type CreateUserStorage interface {
	InsertUser(data *model.UserCreationType) error
}

type createUsersBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUsersBiz {
	return &createUsersBiz{store: store}
}

func (biz *createUsersBiz) CreateNewUser(data *model.UserCreationType) error {
	if err := biz.store.InsertUser(data); err != nil {
		return common.ErrCannotCreateUser(model.EntityName, err)
	}
	return nil
}
