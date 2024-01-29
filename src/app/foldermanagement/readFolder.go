package foldermanagement

import (
	"app/database"
	"errors"
	"fmt"
	"os"
	"time"
)

func ReadFolder(path string) (Sdossier, error, []string){
	files, err := os.ReadDir(path)
	database.Condb()
	currentTime := time.Now()
	output := "success"
	var fileNames []string // Déclarer une slice pour stocker les noms des fichiers

	if path == "" {
		output = "path vide"
		return Sdossier{},errors.New(output), nil
	}
	
	if err != nil {
		output = err.Error()
		
		return Sdossier{}, err, nil
	}

	for _, file := range files {
		fmt.Println(file.Name())
		fileNames = append(fileNames, file.Name()) // Ajouter le nom du fichier à la slice
	}
////////////////////////////////////////
	var children []Children
	for _, childName := range fileNames {
			children = append(children, Children{Name: childName})
	}

	var myFolder = Sdossier{Name: path, Children: children}
	fmt.Println("myFolder",myFolder)
//////////////////////////////////////////

	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "ReadFolder()", path, output)

	return myFolder ,nil, fileNames
}