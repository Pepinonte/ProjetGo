package filesmanagement

import (
	"fmt"
	"os"
)

func DeleteFile(path string) error {
	if path == "" {
		return fmt.Errorf("le nom du fichier ne peut pas être vide")
	}
	err := os.Remove(path)
	if err != nil {

		fmt.Println("impossible de supprimer le fichier")
		return err
	}
	return nil
}