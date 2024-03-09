package usecase

import (
	E "docx-doc-manager-srv/core/entity"
	"docx-doc-manager-srv/core/repo"
	"mime/multipart"
)

type IFSUseCase interface {
	FileUpload(file multipart.File, header *multipart.FileHeader) (string, error)
	FileDownload(documentId string) (E.FSFileObject, error)
}

type FSUseCase struct {
	repo repo.IGridFSRepo
}

func (uc *FSUseCase) FileUpload(file multipart.File, header *multipart.FileHeader) (string, error) {
	documentId, err := uc.repo.GridFSFileUpload(file, header)

	return documentId, err
}

func (uc *FSUseCase) FileDownload(documentId string) (E.FSFileObject, error) {
	object, err := uc.repo.GridFSFileDownload(documentId)

	return object, err
}

func NewFSUseCase(repo repo.IGridFSRepo) IFSUseCase {
	return &FSUseCase{repo: repo}
}
