package foldermanagement

import (
	"app/database"
	"errors"
	"os"
	"time"
)

var ErrPathVoid = errors.New("path vide")

func CreateFolder(path string)  (string, error){
	err := os.Mkdir(path, 0755)
	database.Condb()
	currentTime := time.Now()
	
	if path == ""  {
		return "",errors.New("path vide")
	}
	if err != nil {
		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "CreateFolder()", path, "success")

		return "", err
	}
	return "success", nil
}


