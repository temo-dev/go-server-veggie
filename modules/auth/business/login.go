package business

import (
	commonError "server-veggie/common/error"
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
		return commonError.ErrCannotLogin(model.EntityName, err)
	}
	if password != data.Password {
		return commonError.ErrCannotLogin(model.EntityName, model.ErrWrongPassword)
	} else if password == "" {
		return commonError.ErrCannotLogin(model.EntityName, model.ErrWrongPassword)
	} else {
		return nil
	}
}
