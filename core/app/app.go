package app

import (
	"docx-doc-manager-srv/config"
	v1 "docx-doc-manager-srv/core/controller/rest/v1/router"
	"docx-doc-manager-srv/core/repo"
	"docx-doc-manager-srv/core/usecase"
	"docx-doc-manager-srv/pkg/mongo"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(cfg config.Config) {
	db := mongo.NewMongo(cfg.DB.URI, int64(cfg.DB.Timeout))
	if err := db.Connect(); err != nil {
		os.Exit(1)
	}
	client := db.GetClient()

	uc := usecase.NewFSUseCase(repo.NewGridFS(client, cfg.DB.Database))

	handler := gin.Default()
	v1.NewRouter(handler, uc)

	handler.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
