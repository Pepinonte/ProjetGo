package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func CreateFile(path string) (Sfile, error) {

	database.Condb()
	currentTime := time.Now()
	output := "Success"

	if path == "" {
		output = "le nom du fichier ne peut pas être vide"
		return Sfile{}, fmt.Errorf("le nom du fichier ne peut pas être vide")
	}

	_, err := os.Stat(path)
	if err == nil {
		output = "le fichier existe déjà"
		return Sfile{}, fmt.Errorf("le fichier existe déjà")
	}

	_, err = os.Create(path)
	
	if err != nil {
		output = "impossible de créer le fichier"
		fmt.Println("impossible de créer le fichier")
		return Sfile{}, err
	}

	var myFile = Sfile{Name: path}
		//fmt.Println("myFolder",myFile)
	

	database.AddLog(currentTime.Format("2006-01-02 15:04:05"), "filesmanagement", "createFile", path, output)

	return myFile, nil
}