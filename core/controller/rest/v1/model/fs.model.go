package model

import "mime/multipart"

type FSUploadRequest struct {
	Data *multipart.FileHeader `form:"data" binding:"required"`
}

type FSUploadResponse struct {
	DocumentId   string `json:"documentId" binding:"required,alphanum,len=24"`
	DocumentName string `json:"documentName"`
	Size         int64  `json:"size"`
	Type         string `json:"type"`
}

type FSDownloadRequest struct {
	DocumentId string `uri:"documentId" binding:"required,alphanum,len=24"`
}
