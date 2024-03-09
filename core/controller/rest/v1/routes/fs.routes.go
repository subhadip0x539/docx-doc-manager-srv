package routes

import (
	M "docx-doc-manager-srv/core/controller/rest/v1/model"
	"docx-doc-manager-srv/core/usecase"

	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FSRoutes struct {
	usecase usecase.IFSUseCase
}

func (r *FSRoutes) Upload(c *gin.Context) {
	var request M.FSUploadRequest

	if err := c.ShouldBind(&request); err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	file, err := request.Data.Open()
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	defer file.Close()

	documentId, err := r.usecase.FileUpload(file, request.Data)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, M.FSUploadResponse{
		DocumentId:   documentId,
		DocumentName: request.Data.Filename,
		Size:         request.Data.Size,
		Type:         request.Data.Header.Get("Content-Type"),
	})
}

func (r *FSRoutes) Download(c *gin.Context) {
	var request M.FSDownloadRequest

	if err := c.ShouldBindUri(&request); err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	document, err := r.usecase.FileDownload(request.DocumentId)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+document.DocumentName)
	c.Data(http.StatusOK, document.Type, document.Content)
}

func NewFSRoutes(rg *gin.RouterGroup, uc usecase.IFSUseCase) {
	routes := &FSRoutes{usecase: uc}

	handler := rg.Group("/fs")
	{
		handler.POST("/upload", routes.Upload)
		handler.GET("/download/:documentId", routes.Download)
	}
}
