package foldermanagement

import (
	"errors"
	"os"
)

var ErrPathVoid = errors.New("path vide")

func CreateFolder(path string)  (string, error){
	err := os.Mkdir(path, 0755)
	if path == ""  {
		return "",errors.New("path vide")
	}
	if err != nil {
		return "",err
	}
	return "Dossier créé", nil
}