package utils

import (
	"log"
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
)

func CreateDirectories(directoryNames []string) {

	baseDirectory, err := os.Getwd()
	if err != nil {
		slog.Error(err.Error())
	}

	for _, directoryName := range directoryNames {
		directoryPath := filepath.Join(baseDirectory, directoryName)
		if _, error := os.Stat(directoryPath); os.IsNotExist(error) {
			err := os.MkdirAll(directoryPath, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
