package foldermanagement

import (
	"app/database"
	"errors"
	"fmt"
	"os"
	"time"
)

func ReadFolder(path string) (string, error){
	files, err := os.ReadDir(path)
	database.Condb()
	currentTime := time.Now()
	output := "success"

	if path == "" {
		output = "path vide"
		return "",errors.New(output)
	}
	
	if err != nil {
		output = err.Error()
		
		return "", err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "ReadFolder()", path, output)

	return "succes" ,nil
}