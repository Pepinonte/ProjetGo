package filesmanagement

import (
	"app/database"
	"fmt"
	"os"
	"time"
)

func DeleteFile(path string) (Sfile, error) {
	database.Condb()
	currentTime := time.Now()
	output := "Success"
	if path == "" {
		output = "le nom du fichier ne peut pas être vide"
		return Sfile{}, fmt.Errorf("le nom du fichier ne peut pas être vide")
	}
	err := os.Remove(path)
	if err != nil {
		output = "impossible de supprimer le fichier"
		fmt.Println("impossible de supprimer le fichier")
		return Sfile{}, err
	}

	var myFile = Sfile{Name: path}



	database.AddLog(currentTime.Format("2006-01-02 15:04:05"), "filesmanagement", "deleteFile", path, output)
		
	return myFile,nil
}