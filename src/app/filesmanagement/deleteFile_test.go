package filesmanagement

import (
	"os"
	"testing"
)

func TestDeleteFile(t *testing.T) {
	file, _ := os.Create("gravity.txt")
	file.Close()
	
	err := DeleteFile("gravity.txt")
	if err != nil {
		
		t.Errorf("Error occurred: %v", err)
	}
	
}

func TestDeleteFileWithoutPath(t *testing.T) {
	// delete file without argument path
	err := DeleteFile("")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}


func TestDeleteFileWithNonExistingFile(t *testing.T) {
	// delete file with non existing file
	err := DeleteFile("nonexistingfile.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}
