package router

import (
	"docx-doc-manager-srv/core/controller/rest/v1/routes"
	"docx-doc-manager-srv/core/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, uc usecase.IFSUseCase) {
	handler.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	rg := handler.Group("/v1")
	{
		routes.NewFSRoutes(rg, uc)
	}
}
