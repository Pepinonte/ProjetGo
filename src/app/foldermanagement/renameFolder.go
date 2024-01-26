package foldermanagement

import (
	"errors"
	"os"
)

func RenameFolder(path string, newPath string) (string, error) {
	if path == "" {
		return "", errors.New("path vide")
	}
		err:=os.Rename(path, newPath)

		if err != nil {
			return "", err
		}
		return "succes", nil

}