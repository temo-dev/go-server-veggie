package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/auth/model"
)

type LoginStorage interface {
	SelectUser(data *model.LoginInput) (user *schema.User, err error)
}

type loginBiz struct {
	store LoginStorage
}

func NewLoginBiz(store LoginStorage) *loginBiz {
	return &loginBiz{store: store}
}

func (biz *loginBiz) LoginUser(data *model.LoginInput) (user *schema.User, err error) {
	//handle logic bussiness
	user, err = biz.store.SelectUser(data)
	if err != nil {
		return user, commonError.ErrCannotLogin(model.EntityName, err)
	}
	if user.Password != data.Password {
		return nil, commonError.ErrCannotLogin(model.EntityName, model.ErrWrongPassword)
	} else if user.Password == "" {
		return nil, commonError.ErrCannotLogin(model.EntityName, model.ErrWrongPassword)
	} else {
		return user, nil
	}
}
