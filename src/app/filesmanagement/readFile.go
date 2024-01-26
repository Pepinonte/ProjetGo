package filesmanagement

import (
	"fmt"
	"os"
)

func ReadFile(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("le nom du fichier ne peut pas être vide")
	}

	file, error := os.Open(path)
	if error != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return "", error
	}

	defer file.Close()

	stat, error := file.Stat()
	if error != nil {
		fmt.Println("Impossible d'obtenir les informations sur le fichier")
		return "", error
	}

	bs := make([]byte, stat.Size())
	_, error = file.Read(bs)
	if error != nil {
		fmt.Println("Impossible de lire le fichier")
		return "", error
	}

	str := string(bs)
	fmt.Println(str)

	return "succes", nil
}