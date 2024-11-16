package business

import (
	"server-veggie/common"
	"server-veggie/modules/auth/model"
)

type LoginStorage interface {
	SelectUser(data *model.LoginInput) (password string, err error)
}

type loginBiz struct {
	store LoginStorage
}

func NewLoginBiz(store LoginStorage) *loginBiz {
	return &loginBiz{store: store}
}

func (biz *loginBiz) LoginUser(data *model.LoginInput) (err error) {
	//handle logic bussiness
	password, err := biz.store.SelectUser(data)
	if err != nil {
		return common.ErrCannotLogin(model.EntityName, err)
	}
	if password != data.Password {
		return common.ErrCannotLogin(model.EntityName, model.ErrWrongPassword)
	} else {
		return nil
	}
}
