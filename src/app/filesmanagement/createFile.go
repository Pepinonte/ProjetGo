package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func CreateFile(path string) (string, error) {
	database.Condb()
	currentTime := time.Now()
	output := "success"
	if path == "" {
		output = "path vide"
		return "", fmt.Errorf("le nom du fichier ne peut pas être vide")

	}

	_, err := os.Stat(path)
	if err == nil {
		output = "le fichier existe déjà"
		return "", fmt.Errorf("le fichier existe déjà")
	}

	_, err = os.Create(path)
	if err != nil {
		output = "impossible de créer le fichier"
		fmt.Println("impossible de créer le fichier")
		return "", err
	}

	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filesmanagement", "CreateFile()", path, output)

	return "succes", nil
}