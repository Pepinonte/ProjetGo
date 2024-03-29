package foldermanagement

import (
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	path := "testCreateFolder"
	_,err,_ := CreateFolder(path)

	if err != nil {
		t.Errorf("Une erreur inattendu est survenue lors de l'execution CreateFolder() : %s", err)
	}
	os.Remove(path)
}
func TestCreateFolderWithEmptyPath(t *testing.T) {
	path := ""
	_,err,_ := CreateFolder(path)

	if err == nil {
		t.Errorf("Erreur non retournée lors de l'execution de CreateFolder() avec path vide")
	}
}

// 	path := ""
// 	createFile(path)
// 	err := ReadFolder(path)
// 	if err == nil{
// 		t.Errorf("ReadFolder() did not return an error")
// 	}
// }