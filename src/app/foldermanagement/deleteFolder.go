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
	output := "success"

	if path == ""  {
		output = "path vide"
		return "", errors.New(output)
	}
	if err != nil {
		output = err.Error()
		return "", err
	}

	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "DeleteFolder()", path, output)

	return output, nil
}
