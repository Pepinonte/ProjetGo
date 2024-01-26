package foldermanagement

import (
	"errors"
	"os"
)

func DeleteFolder(path string) (string, error){
	err := os.RemoveAll(path)
	if path == ""  {
		return "", errors.New("path vide")
	}
	if err != nil {
		return "", err
	}
	return "succes", nil
}
