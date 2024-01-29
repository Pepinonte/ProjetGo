package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func ModifyFile(name string, content string) (string, error){

	database.Condb()
	currentTime := time.Now()
	output := "Success"

	if name == "" {
		output = "le nom du fichier ne peut pas être vide"
		return "", fmt.Errorf("le nom du fichier est vide")
	}

	if content == "" {
		output = "le contenu pour le fichier ne peut pas être vide"
		return "", fmt.Errorf("le contenu pour le fichier est vide")
	}

	file, error := os.OpenFile(name, os.O_RDWR|os.O_TRUNC, 0644)
	if error != nil {
		output = "Impossible d'ouvrir le fichier"
		fmt.Println("Impossible d'ouvrir le fichier")
		return "", error
	}

	defer file.Close()

	file.WriteString(content)

	database.AddLog(currentTime.Format("2006-01-02 15:04:05"), "filesmanagement", "modifyFile", name, output)

	return "succes", nil
}