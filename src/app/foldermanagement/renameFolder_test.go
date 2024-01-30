package foldermanagement

import (
	"os"
	"testing"
)

func TestRenameFolder(t *testing.T) {
	path := "testRenameFolder"
	newPath := "testNameRenameFolder"
	os.Mkdir(path, 0755)

	_,err, _:=RenameFolder(path, newPath)

	if err != nil {
		t.Errorf("Une erreur inatendu est survenue lors de l'execution de RenameFolder() : %s", err)
	}
	os.RemoveAll(newPath)
}

func TestRenameFolderWithEmptyPath(t *testing.T) {
	path := ""
	newPath := ""
	os.Mkdir(path, 0755)
	_,err, _:=RenameFolder(path, newPath)

	if err == nil {
		t.Errorf("Erreur retournée lors de l'execution de RenameFolder() avec path vide")
	}
	os.RemoveAll(path)
}