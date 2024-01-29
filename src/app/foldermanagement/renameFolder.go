package foldermanagement

import (
	"app/database"
	"errors"
	"os"
	"time"
)

func RenameFolder(path string, newPath string) (string, error) {
	if path == "" {
		return "", errors.New("path vide")
	}
		err:=os.Rename(path, newPath)
		database.Condb()
	currentTime := time.Now()

		if err != nil {
			database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "RenameFolder()", path, "success")
			
			return "", err
		}
		return "succes", nil

}