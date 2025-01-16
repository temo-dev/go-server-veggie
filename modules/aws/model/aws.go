package model

import "errors"

const (
	EntityName = "aws"
)

var (
	ErrorCreateLinkS3 = errors.New("create link s3 failed")
)

type CreateLinkFileType struct {
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
}
