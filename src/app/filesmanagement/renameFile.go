package filesmanagement

import (
	"fmt"
	"os"
)

func RenameFile(name string, newName string) error {
	if name == "" {
		return fmt.Errorf("le nom du fichier est vide")
	}

	if newName == "" {
		return fmt.Errorf("le nouveau nom du fichier est vide")
	}
	err := os.Rename(name, newName)
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return err
	}

	return nil
}