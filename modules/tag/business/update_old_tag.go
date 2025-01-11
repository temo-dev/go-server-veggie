package business

import (
	commonError "server-veggie/common/error"
	"server-veggie/modules/tag/model"
)

type UpdateOldTagStorage interface {
	UpdateOldTagSQL(cond map[string]interface{}) error
	SelectTagById(cond map[string]interface{}) (*model.TagType, error)
}
type updateOldTagBiz struct {
	store UpdateOldTagStorage
}

func NewUpdateOldTagBiz(store UpdateOldTagStorage) *updateOldTagBiz {
	return &updateOldTagBiz{store: store}
}

func (biz *updateOldTagBiz) UpdateOldTag(newTag *model.TagType) (*model.TagType, error) {
	// check tag exist
	oldTag, err := biz.store.SelectTagById(map[string]interface{}{"tag_id": newTag.TagID})
	if err != nil {
		return nil, commonError.ErrCannotUpdateTag(model.EntityName, err)
	}
	dataUpdate := getUpdatedFields(oldTag, newTag)
	if len(dataUpdate) <= 1 {
		return nil, commonError.ErrCannotUpdateTag(model.EntityName, model.ErrorUpdateTagIsNotChanged)
	}
	//update tag
	if err := biz.store.UpdateOldTagSQL(dataUpdate); err != nil {
		return nil, commonError.ErrCannotUpdateTag(model.EntityName, err)
	}
	//get new tag is updated
	newTagUpdated, err := biz.store.SelectTagById(map[string]interface{}{"tag_id": newTag.TagID})
	if err != nil {
		return nil, commonError.ErrCannotUpdateTag(model.EntityName, err)
	}
	return newTagUpdated, nil
}

func getUpdatedFields(oldTag *model.TagType, newTag *model.TagType) map[string]interface{} {
	updates := make(map[string]interface{})
	if oldTag.TagID == newTag.TagID {
		updates["TagID"] = newTag.TagID
	}
	if oldTag.TagName != newTag.TagName {
		updates["TagName"] = newTag.TagName
	}
	if oldTag.Description != newTag.Description {
		updates["Description"] = newTag.Description
	}
	if oldTag.ImageURL != newTag.ImageURL {
		updates["ImageURL"] = newTag.ImageURL
	}
	return updates
}
