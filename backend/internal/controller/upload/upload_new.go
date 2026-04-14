package upload

import "cozeos/api/upload"

type ControllerV1 struct{}

func NewV1() upload.IUploadV1 {
	return &ControllerV1{}
}
