package foldermanagement

import (
	"app/database"
	"errors"
	"os"
	"time"
)

func DeleteFolder(path string) (string, error){
	err := os.RemoveAll(path)
	database.Condb()
	currentTime := time.Now()

	if path == ""  {
		return "", errors.New("path vide")
	}
	if err != nil {
		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "DeleteFolder()", path, "success")
		
		return "", err
	}
	return "succes", nil
}
