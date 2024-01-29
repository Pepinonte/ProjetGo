package filesmanagement

import (
	"os"
	"testing"
)

func TestModifyFile(t *testing.T) {
	file, _ := os.Create("fichiermodif.txt")

	file.Close()

	err:=ModifyFile("fichiermodif.txt", "test")

	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	
	os.Remove("fichiermodif.txt")
}

func TestModifyFileWithWrongPath(t *testing.T) {
	// modify file with wrong path
	err := ModifyFile("wrongpath/fichiermodif.txt", "test")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestModifyFileWithoutPath(t *testing.T) {
	// modify file without argument path
	err := ModifyFile("", "test")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestModifyFileWithEmptyContent(t *testing.T) {
	// modify file with empty content
	file, _ := os.Create("fichiermodif.txt")
	file.Close()
	err := ModifyFile("fichiermodif.txt", "")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}
	os.Remove("fichiermodif.txt")
}



