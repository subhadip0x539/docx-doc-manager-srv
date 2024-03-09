package main

import (
	"docx-doc-manager-srv/config"
	"docx-doc-manager-srv/core/app"
	"docx-doc-manager-srv/pkg/logger"
	"docx-doc-manager-srv/pkg/motd"
	"log/slog"
	"os"
)

func init() {
	motd.Info()
	logger.InitializeLogger()
}

func main() {
	if err := config.Register(".env", "env", os.Getenv("GIN_MODE")); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	config := config.GetConfig()

	app.Run(config)
}
