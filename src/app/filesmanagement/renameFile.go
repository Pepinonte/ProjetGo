package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func RenameFile(name string, newName string) (string, error) {
	database.Condb()
	currentTime := time.Now()
	output := "Success"
	if name == "" {
		output = "le nom du fichier ne peut pas être vide"
		return "", fmt.Errorf("le nom du fichier est vide")
	}

	if newName == "" {
		output = "le nouveau nom du fichier ne peut pas être vide"
		return "", fmt.Errorf("le nouveau nom du fichier est vide")
	}
	err := os.Rename(name, newName)
	if err != nil {
		output = "Impossible d'ouvrir le fichier"
		fmt.Println("Impossible d'ouvrir le fichier")
		return "", err
	}

	database.AddLog(currentTime.Format("2006-01-02 15:04:05"), "filesmanagement", "renameFile", name, output)

	return "succes", nil
}