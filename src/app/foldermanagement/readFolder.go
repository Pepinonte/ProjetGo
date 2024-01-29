package foldermanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func ReadFolder(path string) (string, error){
	if path == "" {
		path = "."
	}
	files, err := os.ReadDir(path)
	database.Condb()
	currentTime := time.Now()
	
	if err != nil {
		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "ReadFolder()", path, "success")
		return "", err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return "succes" ,nil
}