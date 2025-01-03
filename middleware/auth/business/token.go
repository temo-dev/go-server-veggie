package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/middleware/auth/model"
)

type TokenStrorage interface {
	SelectUser(data string) error
}

type tokenBiz struct {
	store TokenStrorage
}

func NewTokenBiz(store TokenStrorage) *tokenBiz {
	return &tokenBiz{store: store}
}

func (biz *tokenBiz) TokenUser(data string) error {
	if err := biz.store.SelectUser(data); err != nil {
		return commonError.ErrInvalidExpToken(model.EntityName, model.ErrInvalidAccount)
	}
	return nil
}
