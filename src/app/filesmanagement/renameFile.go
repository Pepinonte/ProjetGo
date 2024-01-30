package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func RenameFile(name string, newName string) (Sfile, error) {
	database.Condb()
	currentTime := time.Now()
	output := "Success"
	if name == "" {
		output = "le nom du fichier ne peut pas être vide"
		return Sfile{}, fmt.Errorf("le nom du fichier est vide")
	}

	if newName == "" {
		output = "le nouveau nom du fichier ne peut pas être vide"
		return Sfile{}, fmt.Errorf("le nouveau nom du fichier est vide")
	}
	err := os.Rename(name, newName)
	if err != nil {
		output = "Impossible d'ouvrir le fichier"
		fmt.Println("Impossible d'ouvrir le fichier")
		return Sfile{}, err
	}

	var myFile = Sfile{Name: newName}

	database.AddLog(currentTime.Format("2006-01-02 15:04:05"), "filesmanagement", "renameFile", name, output)

	return myFile, nil
}