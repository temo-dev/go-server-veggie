package business

import (
	"server-veggie/common"
	"server-veggie/modules/user/model"
)

type DeleteUserStorage interface {
	SelectUserById(cond map[string]interface{}) (*model.UserType, error)
	DeleteUserById(cond map[string]interface{}) error
}

type deleteUserBiz struct {
	store DeleteUserStorage
}

func NewDeleteUserBiz(store DeleteUserStorage) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUserById(id int) error {
	//Check user isExist
	user, err := biz.store.SelectUserById(map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetUser(model.EntityName, err)
	}
	//Delete user by id
	if user == nil {
		return common.ErrCannotDeleteUser(model.EntityName, model.ErrorUserNotFound)

	}
	if err := biz.store.DeleteUserById(map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteUser(model.EntityName, err)
	}
	return nil
}
