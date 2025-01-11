package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/tag/model"
)

type FindAllTagStorage interface {
	SelectAllTags() ([]*model.TagType, error)
}
type findAllTagBiz struct {
	store FindAllTagStorage
}

func NewFindAllTagBiz(store FindAllTagStorage) *findAllTagBiz {
	return &findAllTagBiz{store: store}
}

func (biz *findAllTagBiz) FindAllTag() ([]*model.TagType, error) {
	listTags, err := biz.store.SelectAllTags()
	if err != nil {
		return nil, commonError.ErrCannotFindAllTag(model.EntityName, err)
	}
	return listTags, nil
}
