package foldermanagement

import (
	"app/database"
	"errors"
	"fmt"
	"os"
	"time"
)

func DeleteFolder(path string) (Sdossier, error, []string){
	err := os.RemoveAll(path)
	database.Condb()
	currentTime := time.Now()
	output := "success"
	var fileNames []string

	if path == ""  {
		output = "path vide"
		return Sdossier{}, errors.New(output), nil
	}
	if err != nil {
		output = err.Error()
		return Sdossier{}, err, nil
	}

	files, err := os.ReadDir(path)

	for _, file := range files {
		fmt.Println(file.Name())
		fileNames = append(fileNames, file.Name()) // Ajouter le nom du fichier à la slice
	}

	var children []Children
    for _, childName := range fileNames {
        children = append(children, Children{Name: childName})
    }

    var myFolder = Sdossier{Name: path, Children: children}
		fmt.Println("myFolder",myFolder)


	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "DeleteFolder()", path, output)

	return myFolder, nil, fileNames
}
