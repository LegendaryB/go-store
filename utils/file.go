package utils

import (
	"os"
	"path/filepath"
)

func CreateFileIfNotExist(path string) error {
	directory := filepath.Dir(path)

	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if _, err = os.Create(path); err != nil {
			return err
		}
	}

	return nil
}
