package model

import "errors"

const (
	EntityName = "tag"
)

var (
	ErrTagNotFound             = errors.New("tag not found")
	ErrorUpdateTagIsNotChanged = errors.New("update tag is not changed")
)

type TagCreationType struct {
	TagName     string `json:"tag_name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}
type TagType struct {
	TagID       string `json:"tag_id"`
	TagName     string `json:"tag_name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}
