package filesmanagement

import (
	"fmt"
	"os"
)

func CreateFile(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("le nom du fichier ne peut pas être vide")
	}

	_, err := os.Stat(path)
	if err == nil {
		return "", fmt.Errorf("le fichier existe déjà")
	}

	_, err = os.Create(path)
	if err != nil {
		fmt.Println("impossible de créer le fichier")
		return "", err
	}

	return "succes", nil
}