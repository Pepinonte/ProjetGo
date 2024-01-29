package filesmanagement

import (
	"os"
	"testing"
)

func TestRenameFile(t *testing.T) {
	file, _ := os.Create("fichierarename.txt")
	file.Close()
	
	
 	err :=	RenameFile("fichierarename.txt", "fichierrename.txt")
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	os.Remove("fichierrename.txt")
}

func TestRenameFileWithWrongPath(t *testing.T) {
	// rename file with wrong path
	err := RenameFile("badfile.txt", "goodfile.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestRenameFileWithEmptyPath(t *testing.T) {
	// rename file with empty path
	err := RenameFile("", "goodfile.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestRenameFileWithEmptyNewName(t *testing.T) {
	file, _ := os.Create("fichierarename.txt")
	file.Close()
	// rename file with empty new name
	err := RenameFile("fichierarename.txt", "")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}
	os.Remove("fichierarename.txt")

}
