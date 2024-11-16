package business

import (
	"server-veggie/common"
	"server-veggie/modules/user/model"
)

type FindUserStorage interface {
	SelectUserById(cond map[string]interface{}) (*model.UserType, error)
}

type findUserBiz struct {
	store FindUserStorage
}

func NewFindUserBiz(store FindUserStorage) *findUserBiz {
	return &findUserBiz{store: store}
}

func (biz *findUserBiz) FindUserById(id int) (*model.UserType, error) {
	user, err := biz.store.SelectUserById(map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetUser(model.EntityName, err)
	}
	//handle logic
	if user == nil {
		return nil, common.ErrCannotGetUser(model.EntityName, model.ErrorUserNotFound)
	}
	return user, nil
}
