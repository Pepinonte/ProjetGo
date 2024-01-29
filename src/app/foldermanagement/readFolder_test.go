package foldermanagement

import (
	"os"
	"testing"
)

func TestReadFolder(t *testing.T) {
	path := "testReadFolder"
	os.Mkdir(path, 0755)
	_,err := ReadFolder(path)
	if err != nil {
		t.Errorf("Une erreur inatendu est survenue lors de l'execution CreateFolder() : %s", err)
	}
	os.RemoveAll(path)
}

 func TestReadFolderWithEmptyPath(t *testing.T) {
	path := ""
	os.Mkdir(path, 0755)
	_,err := ReadFolder(path)
	if err == nil{
		t.Errorf(" Erreur retournée lors de l'execution de ReadFolder() avec path vide")
	}
	os.RemoveAll(path)
}