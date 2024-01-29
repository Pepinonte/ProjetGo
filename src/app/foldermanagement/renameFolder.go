package foldermanagement

import (
	"app/database"
	"errors"
	"os"
	"time"
)

func RenameFolder(path string, newPath string) (string, error) {
	
		err:=os.Rename(path, newPath)
		database.Condb()
		currentTime := time.Now()
		output := "success"

		if path == "" {
			output = "path vide"
			return "", errors.New(output)
		}

		if err != nil {
			output = err.Error()

			return "", err
		}

		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "RenameFolder()", path, output)

		return output, nil

}