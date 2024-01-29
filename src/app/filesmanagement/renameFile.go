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
	output := "success"
	if name == "" {
		output = "path vide"
		return "", fmt.Errorf("le nom du fichier est vide")
	}

	if newName == "" {
		output = "nouveau nom vide"
		return "", fmt.Errorf("le nouveau nom du fichier est vide")
	}
	err := os.Rename(name, newName)
	if err != nil {
		output = "impossible de renommer le fichier"
		fmt.Println("Impossible d'ouvrir le fichier")
		return "", err
	}

	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filesmanagement", "RenameFile()", name, output)

	return "succes", nil
}