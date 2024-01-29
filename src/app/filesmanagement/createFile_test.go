package filesmanagement

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	
	err := CreateFile("filecreated.txt")
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	os.Remove("filecreated.txt")

}


func TestCreateFileWithWrongPath(t *testing.T) {
	// create file with wrong path
	err := CreateFile("wrongpath/filecreated.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestCreateFileWithExistingFile(t *testing.T) {
	// create file with existing file
	file, _ := os.Create("test.txt")
	file.Close()
	err := CreateFile("test.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}
	os.Remove("test.txt")
}

func TestCreateFileWithEmptyPath(t *testing.T) {
	// create file with empty path
	err := CreateFile("")

	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}
}

