package storage

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

func GetDataDirectory() (string, error) {
	return xdg.DataFile("tuido")
}

func GetDBPath() (string, error) {
	dataDir, err := GetDataDirectory()

	if err != nil {
		return "", err
	}

	return filepath.Join(dataDir, "tuido.db"), nil
}
