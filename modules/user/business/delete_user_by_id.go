package business

import (
	commonError "server-veggie/common/error"
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

func (biz *deleteUserBiz) DeleteUserById(id string) error {
	//Check user isExist
	user, err := biz.store.SelectUserById(map[string]interface{}{"id": id})
	if err != nil {
		return commonError.ErrCannotGetUser(model.EntityName, err)
	}
	//Delete user by id
	if user == nil {
		return commonError.ErrCannotDeleteUser(model.EntityName, model.ErrorUserNotFound)

	}
	if err := biz.store.DeleteUserById(map[string]interface{}{"id": id}); err != nil {
		return commonError.ErrCannotDeleteUser(model.EntityName, err)
	}
	return nil
}
