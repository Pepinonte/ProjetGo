package filesmanagement

import (
	"os"
	"testing"
)

func TestRenameFile(t *testing.T) {
	file, _ := os.Create("fichierarename.txt")
	file.Close()
	
	
 	_,err :=	RenameFile("fichierarename.txt", "fichierrename.txt")
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	os.Remove("fichierrename.txt")
}

func TestRenameFileWithWrongPath(t *testing.T) {
	// rename file with wrong path
	_,err := RenameFile("badfile.txt", "goodfile.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestRenameFileWithEmptyPath(t *testing.T) {
	// rename file with empty path
	_,err := RenameFile("", "goodfile.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestRenameFileWithEmptyNewName(t *testing.T) {
	file, _ := os.Create("fichierarename.txt")
	file.Close()
	// rename file with empty new name
	_,err := RenameFile("fichierarename.txt", "")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}
	os.Remove("fichierarename.txt")

}
