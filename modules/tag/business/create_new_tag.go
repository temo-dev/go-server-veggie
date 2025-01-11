package business

import "server-veggie/modules/tag/model"

type CreateNewTagStorage interface {
	InsertNewTag(data *model.TagCreationType) error
}
type createNewTagBiz struct {
	store CreateNewTagStorage
}

func NewCreateNewTagBiz(store CreateNewTagStorage) *createNewTagBiz {
	return &createNewTagBiz{store: store}
}

func (biz *createNewTagBiz) CreateNewTag(data *model.TagCreationType) error {
	if err := biz.store.InsertNewTag(data); err != nil {
		return err
	}
	return nil
}
