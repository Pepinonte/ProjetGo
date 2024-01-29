package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func DeleteFile(path string) (string, error) {

	database.Condb()
	currentTime := time.Now()
	output := "success"

	if path == "" {
		output = "path vide"
		return "", fmt.Errorf("le nom du fichier ne peut pas être vide")
	}
	err := os.Remove(path)
	if err != nil {
		output = "impossible de supprimer le fichier"

		fmt.Println("impossible de supprimer le fichier")
		return "", err
	}

	database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filesmanagement", "DeleteFile()", path, output)

	return "succes",nil
}