package business

import "server-veggie/modules/auth/model"

type LogoutStorage interface {
}

type logoutBiz struct {
	store LogoutStorage
}

func NewLogoutBiz(store LogoutStorage) *logoutBiz {
	return &logoutBiz{store: store}
}

func (biz *logoutBiz) LogoutUser(data *model.LogoutInput) error {
	return nil
}
