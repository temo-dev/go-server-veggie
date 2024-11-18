package business

import (
	commonError "server-veggie/common/error"
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
		return commonError.ErrCannotCreateUser(model.EntityName, err)
	}
	return nil
}
