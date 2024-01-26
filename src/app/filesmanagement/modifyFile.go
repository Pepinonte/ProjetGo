package filesmanagement

import (
	"fmt"
	"os"
)

func ModifyFile(name string, content string) (string, error){
	if name == "" {
		return "", fmt.Errorf("le nom du fichier est vide")
	}

	if content == "" {
		
		return "", fmt.Errorf("le contenu pour le fichier est vide")
	}

	file, error := os.OpenFile(name, os.O_RDWR|os.O_TRUNC, 0644)
	if error != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return "", error
	}

	defer file.Close()

	file.WriteString(content)

	return "succes", nil
}