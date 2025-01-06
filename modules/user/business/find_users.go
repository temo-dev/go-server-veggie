package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/user/model"
)

type FindUsersStorange interface {
	SelectUsers() ([]*model.UserType, error)
}

type findUsersBiz struct {
	store FindUsersStorange
}

func NewFindUsersBiz(store FindUsersStorange) *findUsersBiz {
	return &findUsersBiz{store: store}
}

func (biz *findUsersBiz) FindUsers() ([]*model.UserType, error) {
	//hanlde business
	users, err := biz.store.SelectUsers()
	if err != nil {
		return nil, commonError.ErrCannotGetUser(model.EntityName, err)
	}
	return users, nil
}
