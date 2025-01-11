package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/tag/model"
)

type DeleteTagByIdStorage interface {
	DeleteTagById(cond map[string]interface{}) error
	SelectTagById(cond map[string]interface{}) (*model.TagType, error)
}
type deleteTagByIdBiz struct {
	store DeleteTagByIdStorage
}

func NewDeleteTagByIdBiz(store DeleteTagByIdStorage) *deleteTagByIdBiz {
	return &deleteTagByIdBiz{store: store}
}

func (biz *deleteTagByIdBiz) DeleteTagById(id string) error {
	//check tag exist
	tag, err := biz.store.SelectTagById(map[string]interface{}{"tag_id": id})
	if err != nil {
		return commonError.ErrCannotDeleteTag(model.EntityName, err)
	}
	if tag == nil {
		return commonError.ErrCannotDeleteTag(model.EntityName, model.ErrTagNotFound)
	}
	if err := biz.store.DeleteTagById(map[string]interface{}{"tag_id": id}); err != nil {
		return commonError.ErrCannotDeleteTag(model.EntityName, err)
	}
	return nil
}
