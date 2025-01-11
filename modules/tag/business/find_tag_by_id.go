package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/tag/model"
)

type FindTagByIdStorage interface {
	SelectTagById(cond map[string]interface{}) (*model.TagType, error)
}

type findTagByIdBiz struct {
	store FindTagByIdStorage
}

func NewFindTagByIdBiz(store FindTagByIdStorage) *findTagByIdBiz {
	return &findTagByIdBiz{store: store}
}

func (biz *findTagByIdBiz) FindTagById(id string) (*model.TagType, error) {
	tag, err := biz.store.SelectTagById(map[string]interface{}{"tag_id": id})
	if err != nil {
		return nil, commonError.ErrCannotFindTagById(model.EntityName, err)
	}
	return tag, nil
}
