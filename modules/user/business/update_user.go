package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/user/model"
)

type UpdateUserStorage interface {
	SelectUserById(cond map[string]interface{}) (*model.UserType, error)
	UpdateUserById(data *model.UpdateUserType) error
}

type updateUserBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUser(data *model.UpdateUserType) error {
	//check isExist user
	user, err := biz.store.SelectUserById(map[string]interface{}{"id": data.Id})
	if err != nil {
		return commonError.ErrCannotGetUser(model.EntityName, err)
	}
	// update user
	if user == nil {
		return commonError.ErrCannotUpdateUser(model.EntityName, model.ErrorUpdateUser)
	}

	if err := biz.store.UpdateUserById(data); err != nil {
		return commonError.ErrCannotUpdateUser(model.EntityName, err)
	}
	return nil
}
