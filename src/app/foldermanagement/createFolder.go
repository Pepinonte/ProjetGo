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
	output := "success"
	
	if path == ""  {
		output = "path vide"
		return "",errors.New(output)
	}
	if err != nil {
		output = err.Error()
		return "", err
	}	

	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "CreateFolder()", path, output)

	return output, nil
}


