package foldermanagement

import (
	"os"
	"testing"
)

func TestDeleteFolder(t *testing.T) {
	path := "testDeleteFolder"
	os.Mkdir(path, 0755)
	err := DeleteFolder(path)
	if err != nil {
		t.Errorf("Une erreur inattendu est survenue lors de l'execution CreateFolder() : %s", err)
	}
	
}

func TestDeleteFolderWithEmptyPath(t *testing.T) {
	path := ""
	os.Mkdir(path, 0755)
	err := DeleteFolder(path)
	if err == nil {
		t.Errorf("Une erreur est survenue lors de l'execution avec le path vide de CreateFolder() : %s", err)
	}
}
