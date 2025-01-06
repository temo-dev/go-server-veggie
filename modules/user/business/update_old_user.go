package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/user/model"
)

type UpdateUserStorage interface {
	SelectUserById(cond map[string]interface{}) (*model.UserType, error)
	UpdateUserById(cond map[string]interface{}) error
}

type updateUserBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUserById(newUser *model.UserType) (*model.UserType, error) {
	//check isExist user
	oldUser, err := biz.store.SelectUserById(map[string]interface{}{"user_id": newUser.UserID})
	if err != nil {
		return nil, commonError.ErrCannotGetUser(model.EntityName, err)
	}
	dataUpdate := getUpdatedFields(oldUser, newUser)
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrUpdateUserIsNotChanged(model.EntityName, model.ErrorUpdateUserIsNotChanged)
	}
	// update user
	if err := biz.store.UpdateUserById(dataUpdate); err != nil {
		return nil, commonError.ErrCannotUpdateUser(model.EntityName, err)
	}
	// get new user is updated
	newUserUpdated, err := biz.store.SelectUserById(map[string]interface{}{"user_id": newUser.UserID})
	if err != nil {
		return nil, commonError.ErrCannotGetUser(model.EntityName, err)
	}

	return newUserUpdated, nil
}

func getUpdatedFields(oldCategory *model.UserType, newCategory *model.UserType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldCategory.UserID == newCategory.UserID {
		updates["UserID"] = newCategory.UserID
	}
	if oldCategory.UserName != newCategory.UserName {
		updates["UserName"] = newCategory.UserName
	}
	if oldCategory.Email != newCategory.Email {
		updates["Email"] = newCategory.Email
	}
	if oldCategory.Status != newCategory.Status {
		updates["Status"] = newCategory.Status
	}
	return updates
}
