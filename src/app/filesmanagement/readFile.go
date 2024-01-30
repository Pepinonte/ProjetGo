package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func ReadFile(path string) (string, error, string) {
	
	database.Condb()
	currentTime := time.Now()
	output := "Success"

	if path == "" {
		output = "le nom du fichier ne peut pas être vide"
		return "", fmt.Errorf("le nom du fichier ne peut pas être vide"), ""
	}

	file, error := os.Open(path)
	if error != nil {
		output = "Impossible d'ouvrir le fichier"
		fmt.Println("Impossible d'ouvrir le fichier")
		return "", error, ""
	}

	defer file.Close()

	stat, error := file.Stat()
	if error != nil {
		output = "Impossible d'obtenir les informations sur le fichier"
		fmt.Println("Impossible d'obtenir les informations sur le fichier")
		return "", error, ""
	}

	bs := make([]byte, stat.Size())
	_, error = file.Read(bs)
	if error != nil {
		output = "Impossible de lire le fichier"
		fmt.Println("Impossible de lire le fichier")
		return "", error, ""
	}

	str := string(bs)
	fmt.Println(str)
	

	database.AddLog(currentTime.Format("2006-01-02 15:04:05"), "filesmanagement", "readFile", path, output)

	return "succes", nil, str
}